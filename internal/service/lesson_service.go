package service

import (
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/mapper"
	"LMSBitLab/internal/model"
	"LMSBitLab/internal/repository"

	"github.com/sirupsen/logrus"
)

type LessonService interface {
	Create(input dto.CreateLessonDTO) (dto.LessonResponseDTO, error)
	GetByID(id uint) (dto.LessonResponseDTO, error)
	GetByChapterID(chapterID uint) ([]dto.LessonResponseDTO, error)
	Update(id uint, input dto.UpdateLessonDTO) (dto.LessonResponseDTO, error)
	Delete(id uint) error
}

type lessonService struct {
	repo repository.LessonRepository
}

func NewLessonService(repo repository.LessonRepository) LessonService {
	return &lessonService{repo: repo}
}

func (s *lessonService) Create(input dto.CreateLessonDTO) (dto.LessonResponseDTO, error) {
	logrus.Info("Creating new lesson")

	lesson := &model.Lesson{
		Name:        input.Name,
		Description: input.Description,
		Content:     input.Content,
		Order:       input.Order,
		ChapterID:   input.ChapterID,
	}

	if err := s.repo.Create(lesson); err != nil {
		logrus.Errorf("Не удалось создать занятие: Name=%s, %v", input.Name, err)
		return dto.LessonResponseDTO{}, err
	}

	logrus.Debugf("Lesson created details: ID=%d, Name=%s, Order=%d, ChapterID=%d", lesson.ID, lesson.Name, lesson.Order, lesson.ChapterID)
	logrus.Infof("Занятие успешно создано: ID=%d, Name=%s", lesson.ID, lesson.Name)
	return mapper.ToLessonResponseDTO(lesson), nil
}

func (s *lessonService) GetByID(id uint) (dto.LessonResponseDTO, error) {
	lesson, err := s.repo.GetByID(id)
	if err != nil {
		logrus.Errorf("Не удалось найти урок ID=%d, %v", id, err)
		return dto.LessonResponseDTO{}, err
	}

	return mapper.ToLessonResponseDTO(lesson), nil
}

func (s *lessonService) GetByChapterID(chapterID uint) ([]dto.LessonResponseDTO, error) {
	lessons, err := s.repo.GetAllByChapterID(chapterID)
	if err != nil {
		logrus.Errorf("Не удалось получить уроки главы: ChapterID=%d, %v", chapterID, err)
		return nil, err
	}

	return mapper.ToLessonResponseDTOList(lessons), nil
}

func (s *lessonService) Update(id uint, input dto.UpdateLessonDTO) (dto.LessonResponseDTO, error) {
	logrus.Infof("Updating lesson: ID=%d", id)

	lesson, err := s.repo.GetByID(id)
	if err != nil {
		logrus.Errorf("Не удалось найти урок для обновления: ID=%d, %v", id, err)
		return dto.LessonResponseDTO{}, err
	}

	lesson.Name = input.Name
	lesson.Description = input.Description
	lesson.Content = input.Content
	lesson.Order = input.Order

	if err := s.repo.Update(lesson); err != nil {
		logrus.Errorf("Не удалось обновить урок: Name=%s, %v", lesson.Name, err)
		return dto.LessonResponseDTO{}, err
	}

	logrus.Debugf("Lesson updated details: ID=%d, Name=%s, Order=%d", lesson.ID, lesson.Name, lesson.Order)
	logrus.Infof("Урок успешно обновлен: Name=%s", lesson.Name)
	return mapper.ToLessonResponseDTO(lesson), nil
}

func (s *lessonService) Delete(id uint) error {
	logrus.Infof("Deleting lesson: ID=%d", id)
	err := s.repo.Delete(id)
	if err != nil {
		logrus.Errorf("Не удалось удалить урок: ID=%d, %v", id, err)
		return err
	}
	logrus.Infof("Урок удален: ID=%d", id)
	return nil
}
