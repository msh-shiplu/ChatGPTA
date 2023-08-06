package main

// EffectiveFeedbackWithExample contains the effective instructions with examples to give feedback to CS students.
const EffectiveFeedbackWithExample = `
Instruction for Effective Feedback on Student's Code:
1. Encouragement should be given. Either the student has made progress or they are experiencing difficulty
	* Care should be taken to highlight to the student that it is some set of actions that has brought  improvement in the student's work.
		For example : You can add “Almost there! ”, “On the right track.”, “This is the right idea, but..”, “Great Work” with your feedback.
2. Your feedback should not give away the answer
	* The feedback should not tell the programmer exactly how to modify their code to get the right answer. The feedback can provide a general strategy but not a specific code correction.
		For example : “Use == not =”, “Declare i here”
3. The feedback should identifies a next action
	* The feedback should be clear and concise. It should be phrased so that the programmer knows what to do next.
		For example : “use the while syntax here”, “Find the reminder”, “Initialize X”
4. Write Complete Sentences as feedback
	* The feedback should be written in a full sentence, including a verb. (The feedback does not need to include punctuation.) There may be an implied subject but it should be the student or refer to a portion of the code.
		For Example: “write this as grade >= 70”, “line 18 and 19 are in the wrong indentation level, try moving them up once”
5. Determine where the students' understanding is lacking (Identify the gap)
	* The feedback should point out how “what the programmer did” is different from “what the answer should do”. You do not have to say what the program should have done explicitly, but must explicitly say what is the problem in the student's code. You are not  just telling this line is wrong, but telling WHY this line is wrong.
		For example: “this loop would include *all* array elements in the total... you want only the odd elements”
6. Use examples to allow students to test their logic
	* Use variable value (variable may be inferred) or example output to allow students to test their logic.
		For example: “What would happen if x was 27?”, “If x was 27, the output would be false”, “This loop runs 999 times”
`

// EffectiveFeedbackWithoutExample contains the effective instructions to give feedback to CS students.
const EffectiveFeedbackWithoutExample = `
Instruction for Effective Feedback on Student's Code:
1. Encouragement should be given. Either the student has made progress or they are experiencing difficulty
	* Care should be taken to highlight to the student that it is some set of actions that has brought  improvement in the student's work.
2. Your feedback should not give away the answer
	* The feedback should not tell the programmer exactly how to modify their code to get the right answer. The feedback can provide a general strategy but not a specific code correction.
3. The feedback should identifies a next action
	* The feedback should be clear and concise. It should be phrased so that the programmer knows what to do next.
4. Write Complete Sentences as feedback
	* The feedback should be written in a full sentence, including a verb. (The feedback does not need to include punctuation.) There may be an implied subject but it should be the student or refer to a portion of the code.
5. Determine where the students' understanding is lacking (Identify the gap)
	* The feedback should point out how “what the programmer did” is different from “what the answer should do”. You do not have to say what the program should have done explicitly, but must explicitly say what is the problem in the student's code. You are not  just telling this line is wrong, but telling WHY this line is wrong.
6. Use examples to allow students to test their logic
	* Use variable value (variable may be inferred) or example output to allow students to test their logic.
`
