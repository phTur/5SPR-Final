package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // метров в километре
	minInH                     = 60   // минут в часе
	stepLengthCoefficient      = 0.45 // длина шага от роста
	walkingCaloriesCoefficient = 0.5  // поправка для ходьбы
)

// Distance считает дистанцию (км) по шагам и росту (м).
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLenMeters := height * stepLengthCoefficient
	return (float64(steps) * stepLenMeters) / mInKm
}

// MeanSpeed считает среднюю скорость (км/ч).
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 || height <= 0 {
		return 0
	}
	distKm := Distance(steps, height)
	hours := duration.Hours()
	if hours <= 0 {
		return 0
	}
	return distKm / hours
}

// RunningSpentCalories считает калории для бега.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid steps")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight")
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid height")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := (weight * meanSpeed * minutes) / minInH
	return calories, nil
}

// WalkingSpentCalories считает калории для ходьбы.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid steps")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight")
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid height")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := (weight * meanSpeed * minutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
