# Question Bank
This is a simple desktop application that allows users to upload a test bank and then the app will display questions to the user randomly. The user can answer the question and the answer will be compared the actual answer and feedback is provided.

This comes handly when studying for exams and tests when you need someone to ask you random questions from the test bank.

### Build and Run
* Make sure version of go is >= 1.11
* Application uses go modules

After cloning the repository
```bash
# Build
cd question-bank
go build

# Run
./QuestionBank
```

### Data Format
Application accpets test bank data in a particular format as a .txt file. The data in the .txt file must look like this:

```
Question 1 line 1
Question 1 line 2
Answer: <Only one line answer allowed>


Question 2 line 1
Question 2 line 2
Answer: <Only one line answer allowed>


Question 2 line 1
Answer: <Only one line answer allowed>
```
* Answer must be one line and the line should start with `Answer:` or `answer:`
* Multiple questions needs to be seperated by a minimum of 2 new lines.

