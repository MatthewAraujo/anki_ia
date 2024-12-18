package utils

const prompten = `<Context> You are a model for creating academic questions. Your focus is to create multiple-choice questions based on a provided text. The questions should assess the understanding of the text, covering its key concepts. 
The goal is for you to create 10 questions about the provided content and return the answers in JSON format, following the specified structure.

Each question should:

Be clear and directly related to the text.
Have 4 answer options (A, B, C, and D).
Have a correct alternative identified in the "answer_key" field.
</Context>
<Instructions>
Upon receiving the text, create 10 questions in the following JSON format:

[
  {
    "question": "Write the full question here",
    "alternatives": {
      "a": "Write alternative A",
      "b": "Write alternative B",
      "c": "Write alternative C",
      "d": "Write alternative D"
    },
    "right_answer": "a"
  },
  {
    "question": "Write the second full question here",
    "alternatives": {
      "a": "Write alternative A",
      "b": "Write alternative B",
      "c": "Write alternative C",
      "d": "Write alternative D"
    },
    "right_answer": "b"
  }
]
Example:
If the provided text is about the Krebs Cycle, a question could be:

{
  "question": "Which enzyme regulates the reaction that transforms oxaloacetate into citrate?",
  "alternatives": {
    "a": "Citrate synthase.",
    "b": "Malate dehydrogenase.",
    "c": "Fumarase.",
    "d": "Isocitrate dehydrogenase."
  },
  "right_answer": "a"
}
<Additional requirements> - Avoid obvious questions; prioritize those that really test knowledge of the text. - Ensure that the **right_answer** is correct. - Do not provide explanations along with the answers, only the JSON format.
</Instructions>
`

const answeren = `[
  {
    "question": "What is the fate of pyruvate under aerobic conditions after glycolysis?",
    "alternatives": {
      "a": "Transformation into lactate.",
      "b": "Transformation into ethanol.",
      "c": "Oxidative decarboxylation to form acetyl-CoA.",
      "d": "Direct conversion into oxaloacetate."
    },
    "right_answer": "c"
  },
  {
    "question": "What is the primary function of the Krebs cycle?",
    "alternatives": {
      "a": "Produce glucose.",
      "b": "Oxidize acetyl-CoA and produce reduced coenzymes.",
      "c": "Store energy in the form of glycogen.",
      "d": "Convert pyruvate into lactate."
    },
    "right_answer": "b"
  },
  {
    "question": "What type of reaction occurs in the mitochondria involving pyruvate before entering the Krebs cycle?",
    "alternatives": {
      "a": "Carboxylation.",
      "b": "Oxidation.",
      "c": "Oxidative decarboxylation.",
      "d": "Reduction."
    },
    "right_answer": "c"
  },
  {
    "question": "What is the role of anaplerotic reactions in the Krebs cycle?",
    "alternatives": {
      "a": "Generate oxygen for the cycle.",
      "b": "Maintain the concentration of cycle intermediates.",
      "c": "Consume the final products of the cycle.",
      "d": "Promote NAD+ reduction."
    },
    "right_answer": "b"
  },
  {
    "question": "What happens to NADH and FADH2 produced in the Krebs cycle?",
    "alternatives": {
      "a": "They are used directly in ATP synthesis.",
      "b": "They are oxidized in the electron transport chain to generate ATP.",
      "c": "They are converted into CO2.",
      "d": "They are stored in the mitochondria."
    },
    "right_answer": "b"
  },
  {
    "question": "Which of the following enzymes is involved in the transformation of pyruvate into acetyl-CoA?",
    "alternatives": {
      "a": "Citrate synthase.",
      "b": "Isocitrate dehydrogenase.",
      "c": "Pyruvate dehydrogenase.",
      "d": "Fumarase."
    },
    "right_answer": "c"
  },
  {
    "question": "What are the main negative allosteric regulators of pyruvate dehydrogenase?",
    "alternatives": {
      "a": "ATP and succinate.",
      "b": "NADH and acetyl-CoA.",
      "c": "CO2 and ADP.",
      "d": "FADH2 and glucose."
    },
    "right_answer": "b"
  },
  {
    "question": "What function do citrate synthase, isocitrate dehydrogenase, and α-ketoglutarate dehydrogenase perform in the Krebs cycle?",
    "alternatives": {
      "a": "They convert CO2 into NADH.",
      "b": "They regulate the metabolic flow of the cycle.",
      "c": "They are responsible for ATP synthesis.",
      "d": "They store energy in the form of ADP."
    },
    "right_answer": "b"
  },
  {
    "question": "Why is the Krebs cycle classified as an amphibolic cycle?",
    "alternatives": {
      "a": "Only because it generates energy.",
      "b": "Because it participates in both catabolic and anabolic pathways.",
      "c": "Because it occurs in an aquatic environment.",
      "d": "Because it does not involve any oxidation."
    },
    "right_answer": "b"
  },
  {
    "question": "What is the role of pyruvate carboxylase in the Krebs cycle?",
    "alternatives": {
      "a": "Decarboxylate pyruvate into acetyl-CoA.",
      "b": "Reduce NAD+ to NADH.",
      "c": "Carboxylate pyruvate to form oxaloacetate.",
      "d": "Oxidize CO2 to form pyruvate."
    },
    "right_answer": "c"
  }
]`

func GetPromptEn() string {
	return prompten
}
func GetAnswerEn() string {
	return answeren
}
