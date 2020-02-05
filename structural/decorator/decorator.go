package decorator

import (
	"errors"
	"fmt"
)

// IngredientAdd ...
type IngredientAdd interface {
	AddIngredient() (string, error)
}

// PizzaDecorator ...
type PizzaDecorator struct {
	Ingredient IngredientAdd
}

// AddIngredient ...
func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

// Meat ...
type Meat struct {
	Ingredient IngredientAdd
}

// AddIngredient ...
func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingerdient field of the Meat")
	}

	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

// Onion ...
type Onion struct {
	Ingredient IngredientAdd
}

// AddIngredient ...
func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Onion")
	}

	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", s, "onion"), nil
}
