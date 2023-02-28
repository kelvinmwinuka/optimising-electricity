# Minimise delivery emissions

Welcome to the Delivery Emissions live coding exercise for HIVED, the eco-friendly delivery company :truck::recycle:

At HIVED, we are striving to minimize the carbon footprint of our deliveries, and this exercise is similar to the problems we solve here at HIVED.

## Instructions

* :speech_balloon: See this exercise as an interactive session, ask us questions as you would if we were working together
* :ok_hand: Aim to write code in the way you would every day - **you will not be penalised for not completing the exercise**

## Exercise

We'd like you to create an app that puts together routes -with different distances- and vehicles -with different emissions- to find the optimal combination, the one that will emit the least amount of CO2.

### Input

1. `routes.json` This contains an array of different routes. Each route will have an ID and a list of stops. Each stop will have a distance from the previous stop in kilometers.

2. `vehicles.json` This contains an array of vehicles. Each vehicle will have an ID, a maximum range in kilometers, and a CO2 emissions per kilometer.

### Output

* List of optimal vehicle-route pairs
* Total CO2 emissions when using the most polluting vehicle for all routes
* Total CO2 emissions when using the most optimal vehicle-route pairing
