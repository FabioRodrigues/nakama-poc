# Project Overview

This project follows certain design principles and practices to ensure code maintainability, testability, and simplicity. Below are some key points to understand the rationale behind the design decisions made in this project.

## Avoiding Custom Database Tables

In this project, we adhere to the recommendation provided in the [Nakama Server Framework documentation](https://heroiclabs.com/docs/nakama/server-framework/go-runtime/function-reference/) which advises against creating custom database tables. Instead, we leverage Nakama's provided functionality and design our application logic around its features. This approach ensures compatibility, scalability, and minimizes the risk of potential conflicts or issues with future updates to Nakama.

## Separation of Concerns

One of the guiding principles in this project is the separation of concerns. By dividing the codebase into logically distinct modules or components, we make it easier to understand, test, and evolve the code over time. Each module or component is responsible for a specific aspect of the application's functionality, promoting code organization and maintainability.

## Minimal Abstraction Layers

While abstraction layers can sometimes be beneficial for decoupling components and promoting code reusability, the Go community generally discourages excessive abstraction. In this project, we have adopted a pragmatic approach by using as few layers as possible while still maintaining a clear and understandable codebase. This decision aims to strike a balance between simplicity and flexibility, allowing for easier comprehension and maintenance of the code.

## No Mock Library Usage

Due to the nature of the codebase in terms of its size and complexity, we have opted not to use any mock library. Instead, we have focused on writing straightforward and easily understandable test cases that directly interact with the components being tested. This approach helps in keeping the test suite lightweight and maintains clarity, especially in smaller projects where the overhead of using mock libraries may outweigh their benefits.

## Continuous Learning

Given more time, a deeper exploration of the Nakama Server Framework documentation would be beneficial. Frameworks often conceal powerful features and tools that can greatly simplify development tasks and enhance productivity. By investing time in understanding the nuances of Nakama, we can uncover hidden gems and leverage them to make our development process even more efficient and effective.

## Assumptions

This project makes certain assumptions to streamline development and testing processes. Specifically, it assumes the existence of a file named `1.0.1.json` within the `core` folder. This file serves as a fixture or sample data source, facilitating testing through API calls. By establishing such conventions, we aim to simplify the setup and execution of test scenarios, promoting efficiency and consistency in the development workflow.

## Running Tests

To run tests, simply call `make test` in your terminal. Ensure that you have Go version 1.21 or above installed and properly configured in your PATH.

## End-to-End Testing

To perform an end-to-end test, follow these steps:

1. Open Nakama server at [http://127.0.0.1:7351](http://127.0.0.1:7351).
2. Navigate to the `API EXPLORER` menu.
3. From the dropdown menu, select the `file_seeker` endpoint.
4. Pass the following payload to the endpoint:
```json
{
  "version": "1.0.1",
  "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
}
```
5. Execute the request.
Note: Although the endpoint is called FileSeeker, the service also saves data to the database. Ideally, we should discuss a different approach. However, for the sake of this example, it's considered valid.
---
Feel free to reach out if you have any questions or suggestions for further improvements.