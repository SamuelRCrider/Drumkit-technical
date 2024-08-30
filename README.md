# Drumkit Technical

This is a simple Go application that will fetch data from the Motive API and send it to a callback URL.

- On first attempt the AWS Lambda will cold start therefore taking longer to respond.
- On first attempt, if you are not logged in, you will be prompted to do so (both Motive and Pipedream).

## The process

1. I created a gofiber server with two GET endpoints:
   - `/`: here we redirect the user to the Motive API to authorize our app. This redirects to Pipedream and exchanges the code for an access token and then redirects back to our server.
   - `/callback`: this endpoint is hit when the user is redirected back to our server. We use the access token to fetch data from the Motive API and progressively send it to Pipedream, then redirect the user to Pipedream.
2. I created a `request` package that handles the communication with the Motive API and the callback URL. There are five functions in this package:
   - `newMotiveAPI`: creates a new Motive API client
   - `fetchComponent`: fetches a specific component from the Motive API
   - `sendCallback`: sends the fetched data to the callback URL
   - `processComponent`: employs `fetchComponent` and `sendCallback` to fetch and send data for a specific component
   - `HandleRequest`: calls `newMotiveAPI` and then `processComponent` for each component type
3. I then began testing this out locally.
   - Go to Problem 1
4. I then created a new AWS Lambda function, built my Go binary, zipped it up, and uploaded it.
   - Go to Problem 2
5. With that, I wrapped up the technical and pushed everything to GitHub.

## How to run it

1. Go to https://arb5dlfzya64pwndewooxljrtu0gcgrz.lambda-url.us-east-2.on.aws/
2. You will be redirected to Motive. If you aren't logged in, you will be prompted to do so.
3. You will be redirected to the Axle Interview1 Pipedream workflow where the Motive code will be exchanged for an access token.
4. Then, you will be redirected back to the server /callback endpoint.
5. The server will then fetch the components from the Motive API, and progressively send them to Pipedream.
6. You will be redirected to the Pipedream workflow "Crider Data Responses" in the Drumkit Interview Data Responses project.
7. Here you can inspect the components in the live events.

## Problems/Notes

#### Problem 1

- I realized that the callback URL wasn't set up to handle the Motive API component data. So, I created a new Pipedream project called `Drumkit Interview Data Responses` and added a workflow to handle the data.

#### Problem 2

- Initially my Lambda was throwing 500 errors because I hadn't configured my gofiber server to work on AWS Lambda. My first solution here was to set up an API Gateway and refactor my code to work with that. As I tested this solution I realized I was over-engineering it.
- The simplest solution was to add a layer to my Lambda function that could handle the gofiber server while preserving the ability to run it locally.

#### Note 1

- On initial test, user must log into Motive and Pipedream if not already logged in. This is expected behavior.

#### Note 2

- The Axle Interview1 Pipedream workflow was a bit confusing. It already had code written in it and said code seemed to be from a previous candidate. I didn't want to mess with it that much so I just changed the redirect URI to my own and moved on.

## Resources

https://github.com/awslabs/aws-lambda-web-adapter?tab=readme-ov-file
