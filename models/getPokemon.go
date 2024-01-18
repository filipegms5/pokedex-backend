package models

type Pokemon struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Info     string `json:"info"`
	PhotoUrl string `json:"url"`
	Types    []Type `json:"types"`
}

type Type struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

var pokemonList = []Pokemon{
	{
		Id:       1,
		Name:     "Bulbasaur",
		Info:     "Bulbasaur can be seen napping in bright sunlight. There is a seed on its back. By soaking up the sun’s rays, the seed grows progressively larger.",
		PhotoUrl: "https://assets.pokemon.com/assets/cms2/img/pokedex/full/001.png",
		Types:    []Type{{Name: "Grass", Color: "#78C850"}, {Name: "Poison", Color: "#A040A0"}},
	},

	{
		Id:       2,
		Name:     "Ivysaur",
		Info:     "There is a bud on this Pokémon’s back. To support its weight, Ivysaur’s legs and trunk grow thick and strong. If it starts spending more time lying in the sunlight, it’s a sign that the bud will bloom into a large flower soon.",
		PhotoUrl: "https://assets.pokemon.com/assets/cms2/img/pokedex/full/002.png",
		Types:    []Type{{Name: "Grass", Color: "#78C850"}, {Name: "Poison", Color: "#A040A0"}},
	},
	{
		Id:       3,
		Name:     "Venusaur",
		Info:     "There is a large flower on Venusaur’s back. The flower is said to take on vivid colors if it gets plenty of nutrition and sunlight. The flower’s aroma soothes the emotions of people.",
		PhotoUrl: "https://assets.pokemon.com/assets/cms2/img/pokedex/full/003.png",
		Types:    []Type{{Name: "Grass", Color: "#78C850"}, {Name: "Poison", Color: "#A040A0"}},
	},
	{
		Id:       4,
		Name:     "Charmander",
		Info:     "The flame on its tail indicates Charmander’s life force. If it is healthy, the flame burns brightly.",
		PhotoUrl: "https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
		Types:    []Type{{Name: "Fire", Color: "#F08030"}},
	},
	{
		Id:       5,
		Name:     "Charmeleon",
		Info:     "Charmeleon mercilessly destroys its foes using its sharp claws. If it encounters a strong foe, it turns aggressive. In this excited state, the flame at the tip of its tail flares with a bluish white color.",
		PhotoUrl: "https://assets.pokemon.com/assets/cms2/img/pokedex/full/005.png",
		Types:    []Type{{Name: "Fire", Color: "#F08030"}},
	},
}

func FetchAllPokemon() []Pokemon {
	return pokemonList
}
func FetchPokemonById(id int) Pokemon {
	for _, pokemon := range pokemonList {
		if pokemon.Id == id {
			return pokemon
		}
	}
	return Pokemon{}

}
