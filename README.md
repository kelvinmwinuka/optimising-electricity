# Optimising electricity consumption

Welcome to the Delivery Emissions live coding exercise for HIVED, the eco-friendly delivery company :truck::recycle:

At HIVED, we will never use fossil fuels, so all our vans are electric. :zap:
But different sizes of vans will vary in their electricity consumption, and so it's better for the planet (and for our wallet! :money_with_wings:) to carefully select which vans are used for which delivery routes.

## Exercise

We'd like you to write a program that assigns vehicles to routes, in a way that will minimise total energy consumption.

### Input

1. `routes.json` This contains an array of different routes. Each route will have an ID and a list of stops. Each stop will have a distance from the previous stop in kilometers.

2. `vehicles.json` This contains an array of vehicles. Each vehicle will have an ID, a capacity in kWh, and an electricity consumption in kWh/100km.

### Questions

1. Let's establish a baseline: What's the total energy in kWh required to complete all routes using the least efficient vehicle. You can ignore the vehicle's range for this question.
2. Find the list of vehicle-route pairs that will minimise total energy consumption. A vehicle can at most be assigned to one route.
3. Compare the total energy consumed using the optimal vehicle-route pairing to the baseline established in question 1.

## Instructions

### For live sessions

* :speech_balloon: See this exercise as an interactive session, ask us questions as you would if we were working together
* :ok_hand: Aim to write code in the way you would every day - **you will not be penalised if you don't complete the exercise**


### For take-homes

* :pencil2: Feel free to reach out to ask for clarifications and document your assumptions
* :ok_hand: Aim to write code in the way you would every day - we expect you to complete the exercise
