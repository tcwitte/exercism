package lasagna

// Function PreparationTime returns an estimate of the total preparation time based on the provided number of layers and time per layer.
func PreparationTime(layers []string, timePerLayerMinutes int) int {
	if timePerLayerMinutes == 0 {
		timePerLayerMinutes = 2
	}
	return timePerLayerMinutes * len(layers)
}

const noodlesPerLayerGrams = 50
const saucePerLayerLiters = 0.2

// Function Quantities returns the quantity of noodles and sauce needed to prepare the given layers.
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += noodlesPerLayerGrams
		case "sauce":
			sauce += saucePerLayerLiters
		}
	}
	return
}

// Function AddSecretIngredient modifies ownIngredients so the last ingredient matches the friendsIngredients.
func AddSecretIngredient(friendsIngredients, ownIngredients []string) {
	ownIngredients[len(ownIngredients)-1] = friendsIngredients[len(friendsIngredients)-1]
}

// Function ScaleRecipe takes quantities for 2 portions and scales them to the given number of portions.
func ScaleRecipe(quantitiesTwoPortions []float64, portions int) []float64 {
	quantities := make([]float64, len(quantitiesTwoPortions))

	for i := 0; i < len(quantitiesTwoPortions); i++ {
		quantities[i] = float64(portions) * quantitiesTwoPortions[i] / 2
	}

	return quantities
}
