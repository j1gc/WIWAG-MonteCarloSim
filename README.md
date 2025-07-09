# Monte Carlo [WIWAG](https://playeconomy.de/planspiele/wiwag/)

This repository implements a Monte Carlo simulation to forecast key company figures for the business simulation game
[WIWAG](https://playeconomy.de/planspiele/wiwag/), which I played with my team at our school.

[WIWAG](https://playeconomy.de/planspiele/wiwag/) (WIrtschafts-Wochen-Aktien-Gesellschaft) is a business simulation game, provided by the Joachim Herz Stiftung.
During the game, different teams (companies) compete against one another over the span of multiple simulated years.
In these years, the teams implement their strategies by making decisions on their spending,
for example, marketing, R&D, count of produced products, etc.
The goal of the game is to give students a better understanding of business economics.

## Motivation

This project was part of an initiative, started by me, to transform our company motto from Facebook’s original:

> _Move fast and break things_

to Facebook’s newer motto:

> _Move fast with stable infrastructure_

Because just like Facebook, we learned, as our company grew, that stability is as important as the speed of execution.
To have more stability without slowing down, I implemented this Monte Carlo simulation.
It predicts the most likely outcomes of variables such as unit cost,
given uncertain inputs like sales, material consumption, and other factors.
This helped us make data oriented decisions, which ultimately led us to become the most profitable company by far.

![An image showing the simulation viewer](/screenshot_website.png)

## Usage
You can change the simulation behavior inside the runSimulationStep function in backend/routes/run_simulation.go. <br>
Here is an example simulation:

### Example

```` go
func runSimulationStep(input simulation.SimInput) simulation.SimResults {
    currentInput := input
	
	// generates a random normal distributed value in a percentage range under the provided value
	onlineShopInland := RandomNormFloatInPercentageRangeUnder(input.AbsatzmengeInlandOnlineshopHoheQualität, 0.2)
	
	// assigns a new random value to the input value
	currentInput.AbsatzmengeInlandOnlineshopHoheQualität = onlineShopInland
	
	
	simulationData := simulation.SimData{
		Input:   currentInput,
		Results: simulation.SimResults{},
	}
	// runs simulation with specified parameters
	results := simulationData.GetResults()

	//stückSelbstkosten := results.Selbstkosten / (results.AbsatzmengeInlandGesamt + results.AbsatzmengeAuslandGesamt)
	return results
}
````

## ⚠️ Disclaimer ⚠️
This project was also a learning exercise to get more familiar with Golang, SvelteKit, and D3.js.
Because of that, the code is far from perfect, but it still provides business value. <br>
Some thoughts about what I would improve if I restarted this project:

- Build the backend logic in a lower level language like Rust to improve CPU and memory utilization.

- Extract the domain specific simulation logic into nodes that can be edited in the frontend.
  This would make defining simulations more flexible without having to recompile the backend for every change.

- Many variables in the simulation logic are in German and have long names due to the simulation domain.
  By extracting the logic into nodes, I would be able to remove that German from my codebase.

- Since it was a learning project,
  I avoided dependencies for some areas I wanted to understand better, like histogram creation.
  This led to bugs.
  In a rewrite, I would use dependencies like Gonum instead.

## Installation

### Linux etc.
````
git clone https://github.com/j1gc/WIWAG-MonteCarloSim.git

cd WIWAG-MonteCarloSim

// Frontend
cd frontend
bun install
bun run dev

// Backend
cd backend
go run ./cmd

// now open localhost:5173 in your browser and press start
````

### Docker
````
git clone https://github.com/j1gc/WIWAG-MonteCarloSim.git

cd WIWAG-MonteCarloSim

touch ./backend/backend.env

docker compose up

// now open localhost:3000 in your browser and press start
````
