package openai

/*
LATEST MODEL	DESCRIPTION	MAX REQUEST	TRAINING DATA
gpt-3.5-turbo	Most capable GPT-3.5 model and optimized for chat at 1/10th the cost of text-davinci-003. Will be updated with our latest model iteration.	4,096 tokens	Up to Sep 2021
gpt-3.5-turbo-0301	Snapshot of gpt-3.5-turbo from March 1st 2023. Unlike gpt-3.5-turbo, this model will not receive updates, and will only be supported for a three month period ending on June 1st 2023.	4,096 tokens	Up to Sep 2021
text-davinci-003	Can do any language task with better quality, longer output, and consistent instruction-following than the curie, babbage, or ada models. Also supports inserting completions within text.	4,000 tokens	Up to Jun 2021
text-davinci-002	Similar capabilities to text-davinci-003 but trained with supervised fine-tuning instead of reinforcement learning	4,000 tokens	Up to Jun 2021
code-davinci-002	Optimized for code-completion tasks	4,000 tokens	Up to Jun 2021

*/

const (
	ModelGPT3Dot5Turbo     = "gpt-3.5-turbo"
	ModelGPT3Dot5Turbo0301 = "gpt-3.5-turbo-0301"
	ModelTextDavinci003    = "text-davinci-003"
	ModelTextDavinci002    = "text-davinci-002"
	ModelCodeDavinci002    = "code-davinci-002"
)
