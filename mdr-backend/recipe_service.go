package main

type RecipeService struct {
    storage RecipeStorage
}

func (s *RecipeService) ListRecipes() []*Recipe {
    return s.storage.List()
}

func (s *RecipeService) GetRecipe(id string) *Recipe {
    return s.storage.Get(id)
}

func (s *RecipeService) CreateRecipe(recipe Recipe) Recipe {
    s.storage.Create(recipe)
    return recipe
}

func (s *RecipeService) UpdateRecipe(id string, recipe Recipe) *Recipe {
    return s.storage.Update(id, recipe)
}

func (s *RecipeService) DeleteRecipe(id string) *Recipe {
    return s.storage.Delete(id)
}
