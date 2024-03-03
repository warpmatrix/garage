import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

tokenizer = AutoTokenizer.from_pretrained('meta-llama/Llama-2-7b-hf')
model = AutoModelForCausalLM.from_pretrained("meta-llama/Llama-2-7b-hf").half().to(device="cuda:0")
model = torch.compile(model, backend="inductor") # This is the only line of code that we changed
text = "Replace me by any text you'd like."
encoded_input = tokenizer(text, return_tensors='pt').to(device="cuda:0")
output = model(**encoded_input)
print(output)
