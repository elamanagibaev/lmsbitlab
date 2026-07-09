package service

import (
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/mapper"
	"LMSBitLab/internal/model"
	"LMSBitLab/internal/repository"

	"github.com/sirupsen/logrus"
)

type CourseService interface {
	Create(input dto.CreateCourseDTO) (dto.CourseResponseDTO, error)
	GetByID(id uint) (dto.CourseResponseDTO, error)
	GetAll() ([]dto.CourseResponseDTO, error)
	Update(id uint, input dto.UpdateCourseDTO) (dto.CourseResponseDTO, error)
	Delete(id uint) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) Create(input dto.CreateCourseDTO) (dto.CourseResponseDTO, error) {
	course := &model.Course{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := s.repo.Create(course); err != nil {
		logrus.Errorf("не удалось создать курс: %v", err)
		return dto.CourseResponseDTO{}, err
	}

	logrus.Infof("Курс создан: ID:%d, Name=%s", course.ID, course.Name)
	return mapper.ToCourseResponseDTO(course), nil
}

func (s *courseService) GetByID(id uint) (dto.CourseResponseDTO, error) {
	course, err := s.repo.GetByID(id)
	if err != nil {
		return dto.CourseResponseDTO{}, err
	}

	return mapper.ToCourseResponseDTO(course), nil
}

func (s *courseService) GetAll() ([]dto.CourseResponseDTO, error) {
	courses, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return mapper.ToCourseResponseDTOList(courses), nil
}

func (s *courseService) Update(id uint, input dto.UpdateCourseDTO) (dto.CourseResponseDTO, error) {
	course, err := s.repo.GetByID(id)
	if err != nil {
		logrus.Errorf("Не удалось найти курс для обновления: ID:%d, %v", id, err)
		return dto.CourseResponseDTO{}, err
	}

	course.Name = input.Name
	course.Description = input.Description

	if err := s.repo.Update(course); err != nil {
		logrus.Errorf("Не удалось обновить курс: ID=%d, %v", id, err)
		return dto.CourseResponseDTO{}, err
	}

	logrus.Infof("Курс обновлен: ID=%d, Name=%s", course.ID, course.Name)
	return mapper.ToCourseResponseDTO(course), nil
}

func (s *courseService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		logrus.Errorf("Не удалось удалить курс: ID=%d, %v", id, err)
		return err
	}

	logrus.Infof("Курс удален: ID=%d", id)
	return nil
}
