package app

import (
	"careerhub-dataprovider/careerhub/provider/domain/company"
	"careerhub-dataprovider/careerhub/provider/domain/jobposting"
	"careerhub-dataprovider/careerhub/provider/queue"
	"careerhub-dataprovider/careerhub/provider/queue/message_v1"
	"careerhub-dataprovider/careerhub/provider/source"
	"careerhub-dataprovider/careerhub/provider/utils/cchan"
	"fmt"
)

type AppService struct {
	src            source.JobPostingSource
	jobpostingRepo *jobposting.JobPostingRepo
	companyRepo    *company.CompanyRepo
	closedJpQueue  queue.Queue
	companyQueue   queue.Queue
}

func NewAppService(src source.JobPostingSource, repo *jobposting.JobPostingRepo) *AppService {
	return &AppService{
		src:            src,
		jobpostingRepo: repo,
	}
}

func (as *AppService) Run(quitChan <-chan QuitSignal) error {
	savedJpIds, err := as.jobpostingRepo.GetAllHiring(as.src.Site())
	if err != nil {
		return err
	}

	hiringJpIds, err := source.AllJobPostingIds(as.src)
	if err != nil {
		return err
	}

	separatedIds := SeparateIds(savedJpIds, hiringJpIds)

	ProcessClosedJobPostings(as.jobpostingRepo, as.closedJpQueue, separatedIds.ClosePostingIds)

	errChan := make(chan error, 100)
	detailChan := CallDetailApi(as.src, separatedIds.NewPostingIds, errChan, quitChan)
	processedDetailChan := ProcessCompany(as.src, as.companyRepo, as.companyQueue, detailChan, errChan, quitChan)

	processedChan := SendJobPostingInfo(as.jobpostingRepo, as.closedJpQueue, processedDetailChan, errChan, quitChan)

	return nil
}

func CallDetailApi(src source.JobPostingSource, newJpIds []*source.JobPostingId, errChan chan<- error, quitChan <-chan QuitSignal) <-chan *source.JobPostingDetail {
	messageChan := make(chan *source.JobPostingDetail)
	go callDetailApi(src, newJpIds, messageChan, errChan, quitChan)

	return messageChan
}

func callDetailApi(src source.JobPostingSource, newJpIds []*source.JobPostingId, resultChan chan<- *source.JobPostingDetail, errChan chan<- error, quitChan <-chan QuitSignal) {
	defer close(resultChan)

	for _, newJpId := range newJpIds {
		detail, err := src.Detail(newJpId)

		ok := cchan.SendResult(detail, err, resultChan, errChan, quitChan)

		if !ok {
			return
		}
	}
}

func SendJobPostingInfo(jpRepo *jobposting.JobPostingRepo, queue queue.Queue, messageChan <-chan *source.JobPostingDetail, errChan chan<- error, quitChan <-chan QuitSignal) chan<- ProcessedSignal {
	processedChan := make(chan ProcessedSignal, 100)
	go sendJobPostingInfo(jpRepo, queue, messageChan, processedChan, errChan, quitChan)

	return processedChan
}
func sendJobPostingInfo(jpRepo *jobposting.JobPostingRepo, queue queue.Queue, messageChan <-chan *source.JobPostingDetail, processedChan chan<- ProcessedSignal, errChan chan<- error, quitChan <-chan QuitSignal) {
	defer close(processedChan)

	for {
		received, ok := cchan.ReceiveOrQuit(messageChan, quitChan)
		if !ok {
			return
		}

		detail := *received

		message := message_v1.JobPostingInfo{
			Site:        detail.Site,
			PostingId:   detail.PostingId,
			CompanyId:   detail.CompanyId,
			CompanyName: detail.CompanyName,
			JobCategory: detail.JobCategory,
			MainContent: &message_v1.MainContent{
				PostUrl:        detail.MainContent.PostUrl,
				Title:          detail.MainContent.Title,
				Intro:          detail.MainContent.Intro,
				MainTask:       detail.MainContent.MainTask,
				Qualifications: detail.MainContent.Qualifications,
				Preferred:      detail.MainContent.Preferred,
				Benefits:       detail.MainContent.Benefits,
				RecruitProcess: detail.MainContent.RecruitProcess,
			},
			RequiredSkill: detail.RequiredSkill,
			Tags:          detail.Tags,
			RequiredCareer: &message_v1.Career{
				Min: detail.RequiredCareer.Min,
				Max: detail.RequiredCareer.Max,
			},
			PublishedAt: detail.PublishedAt,
			ClosedAt:    detail.ClosedAt,
			Address:     detail.Address,
		}

		//TODO: modify queue interface
		// queue.Send(message)
		_, err := jpRepo.Save(jobposting.NewJobPosting(message.Site, message.PostingId))

		cchan.SendResult(ProcessedSignal{}, err, processedChan, errChan, quitChan)
	}
}

func ProcessClosedJobPostings(jpRepo *jobposting.JobPostingRepo, queue queue.Queue, closedJpIds []*jobposting.JobPostingId) {
	closedJpIdMessages := make([]*message_v1.JobPostingId, len(closedJpIds))

	for i, closedJpId := range closedJpIds {
		closedJpIdMessages[i] = &message_v1.JobPostingId{
			Site:      closedJpId.Site,
			PostingId: closedJpId.PostingId,
		}
	}

	message := message_v1.ClosedJobPostings{
		JobPostingIds: closedJpIdMessages,
	}

	//TODO: modify queue interface
	// queue.Send(message)
	fmt.Println(message)

	//TODO: delete closed job postings
	// jpRepo.DeleteAll(closedJpIds)
}

func ProcessCompany(src source.JobPostingSource,
	companyRepo *company.CompanyRepo, //TODO: need to implement
	queue queue.Queue, detailChan <-chan *source.JobPostingDetail, errChan chan<- error, quitChan <-chan QuitSignal) <-chan *source.JobPostingDetail {

	prosessedChan := make(chan *source.JobPostingDetail)

	go func() {
		for {
			received, ok := cchan.ReceiveOrQuit(detailChan, quitChan)
			if !ok {
				return
			}

			detail := *received

			err := processCompany(src, companyRepo, queue, &company.CompanyId{
				Site:      detail.Site,
				CompanyId: detail.CompanyId,
			})

			cchan.SendResult(detail, err, prosessedChan, errChan, quitChan)
		}
	}()

	return prosessedChan
}

func processCompany(
	src source.JobPostingSource,
	companyRepo *company.CompanyRepo, //TODO: need to implement
	queue queue.Queue, companyId *company.CompanyId) error {

	companyInfo, err := companyRepo.Get(companyId)

	if err != nil {
		//TODO: error handling
		return err
	} else if companyInfo != nil {
		return nil // already processed
	}

	srcCompany, err := src.Company(companyId.CompanyId)

	if err != nil {
		//TODO: error handling
		return err
	}

	message := message_v1.Company{
		Site:          srcCompany.Site,
		CompanyId:     srcCompany.CompanyId,
		Name:          srcCompany.Name,
		CompanyUrl:    srcCompany.CompanyUrl,
		CompanyImages: srcCompany.CompanyImages,
		Description:   srcCompany.Description,
		CompanyLogo:   srcCompany.CompanyLogo,
	}

	//TODO: modify queue interface
	// queue.Send(message)
	fmt.Println(message)

	_, err = companyRepo.Save(company.NewCompany(src.Site(), companyId.CompanyId))

	return err
}
