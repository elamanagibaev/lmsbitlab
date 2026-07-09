package service

import (
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/mapper"
	"LMSBitLab/internal/model"
	"LMSBitLab/internal/repository"

	"github.com/sirupsen/logrus"
)

type ChapterService interface {
	Create(input dto.CreateChapterDTO) (dto.ChapterResponseDTO, error)
	GetByID(id uint) (dto.ChapterResponseDTO, error)
	GetByCourseID(courseID uint) ([]dto.ChapterResponseDTO, error)
	Update(id uint, input dto.UpdateChapterDTO) (dto.ChapterResponseDTO, error)
	Delete(id uint) error
}

type chapterService struct {
	repo repository.ChapterRepository
}

func NewChapterService(repo repository.ChapterRepository) ChapterService {
	return &chapterService{repo: repo}
}

func (s *chapterService) Create(input dto.CreateChapterDTO) (dto.ChapterResponseDTO, error) {
	chapter := &model.Chapter{
		Name:        input.Name,
		Description: input.Description,
		Order:       input.Order,
		CourseID:    input.CourseID,
	}

	if err := s.repo.Create(chapter); err != nil {
		logrus.Errorf("Не удалось создать главу %v", err)
		return dto.ChapterResponseDTO{}, err
	}

	logrus.Infof("Глава успешна создана: ID=%d, Name=%s", chapter.ID, chapter.Name)
	return mapper.ToChapterResponseDTO(chapter), nil
}

func (s *chapterService) GetByID(id uint) (dto.ChapterResponseDTO, error) {
	chapter, err := s.repo.GetByID(id)
	if err != nil {
		return dto.ChapterResponseDTO{}, err
	}

	return mapper.ToChapterResponseDTO(chapter), nil
}

func (s *chapterService) GetByCourseID(courseID uint) ([]dto.ChapterResponseDTO, error) {
	chapters, err := s.repo.GetAllByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	return mapper.ToChapterResponseDTOList(chapters), nil
}

func (s *chapterService) Update(id uint, input dto.UpdateChapterDTO) (dto.ChapterResponseDTO, error) {
	chapter, err := s.repo.GetByID(id)
	if err != nil {
		logrus.Errorf("Не удалось найти главу для обновления, ID=%d, %v", id, err)
		return dto.ChapterResponseDTO{}, err
	}

	chapter.Name = input.Name
	chapter.Description = input.Description
	chapter.Order = input.Order

	if err := s.repo.Update(chapter); err != nil {
		logrus.Errorf("Не удалось обновить главу, Name=%s, %v", chapter.Name, err)
		return dto.ChapterResponseDTO{}, err
	}

	logrus.Infof("Глава успешна обновлена ID=%d, Name=%s", chapter.ID, chapter.Name)
	return mapper.ToChapterResponseDTO(chapter), nil
}

func (s *chapterService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		logrus.Errorf("Не удалось удалить главу, ID=%d, %v", id, err)
		return err
	}

	logrus.Infof("Глава удалена: ID=%d", id)
	return nil
}
