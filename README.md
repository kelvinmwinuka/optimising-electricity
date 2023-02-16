# Minimize CO2 Emissions for Delivery Routes

Welcome to the Delivery Routes pairing exercise for HIVED, the world's most eco-friendly parcel delivery company. At HIVED, we are striving to minimize the carbon footprint of our deliveries, and this exercise is similar to the problems we solve here at HIVED.

## The Problem

You are given two JSON files.

1. `routes.json` This contains an array of different routes. Each route will have an ID and a list of stops. Each stop will have a distance from the previous stop in kilometers.

2. `vehicles.json` This contains an array of vehicles. Each vehicle will have an ID, a maximum range in kilometers, and a CO2 emissions per kilometer.

We'd like you to create an app that goes through these json files and finds the optimal vehicle for each route that results in the lowest total CO2 emissions.

## Expected Output

Output a list of the vehicle IDs and the route they will complete.

Output the total CO2 emissions for the solution.

Take the most inefficient vehicle from the `vehicles.json` array and output the total C02 emissions if only this type of vehicle would complete the routes.

## The Rules

Approach this problem any way you like, you can build a RESTful service, a command line app or a UI if you wish.

We won't spend the whole session writing code; we like to leave plenty of time at the end for you to ask us questions. Don't feel you have to rush a solution. You will not be judged on whether you finish or not. We would much rather you give us a solution that you are proud of and feel best displays the type of code you would want to ship into production. That said, make sure there is a way to run your app and it does give some sort of output.

We'll approach this in a very similar way to how we pair at HIVED day-to-day. We'd like you to lead the session, but we'll be assisting you and asking questions along the way. We're there to help build the app, not to catch you out.