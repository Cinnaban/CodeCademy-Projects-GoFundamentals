package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("You have", fuel, "gallons of fuel left!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
  }
  return fuel
}

// Create the function greetPlanet() here
func greenPlanet(planet string) {
  fmt.Println("Welcome to planet", planet)
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining int
  var fuelCost int
  
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greenPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}



func main() {
  // Test your functions!
  // Create `planetChoice` and `fuel`
  var fuel int
  fuel = 1000000

  var planetChoice string
  planetChoice = "Venus"

  fuel = flyToPlanet(planetChoice, fuel)
  // And then liftoff!
  fuelGauge(fuel)
}