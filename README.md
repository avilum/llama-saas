# llama-saas
A real-time client and server for LLaMA.<br>
- 🚀 Runs on any CPU machine, with no need for GPU 🚀<br>
- The server is written in Go.
- The client is written in Python using requests with response streaming in real time.

I personally used the smallest `7B/` model on an Intel PC / Macbook Pro, which is ~4.8G when quantized to 4 bit, or ~13G in full precision.

# Examples
- Nice example: elaborate about "Github"  
<img src="llama-server-example-1.gif">

- Biased example: elaborate about "Donald Trump"  
<img src="llama-server-example-2.gif">

### Get LLaMA Pretrained Checkpoints
Note that LLaMA <a href="https://ai.facebook.com/blog/large-language-model-llama-meta-ai/">cannot be used for commercial use</a>.
- > To maintain integrity and prevent misuse, we are releasing our model under a noncommercial license focused on research use cases. Access to the model will be granted on a case-by-case basis to academic researchers; those affiliated with organizations in government, civil society, and academia; and industry research laboratories around the world. People interested in applying for access can find the link to the application in our research paper.

Apply for <a href="https://docs.google.com/forms/d/e/1FAIpQLSfqNECQnMkycAp2jP4Z9TFX0cGR4uf7b_fBxjY_OjhJILlKGA/viewform">Official Access</a>. You will get a unique download link once you are approved.
# How to use
Assuming you have the LLaMA checkpoints (☝️)
1. Clone and build https://github.com/ggerganov/llama.cpp
2. Edit the `LLAMA_MODEL_PATH` and `LLAMA_MAIN` variables in `server.go`.
3. Build and run the server: 
```shell
go build
./server
```
4. Run the client:
```shell
python3 -m pip install requests
python3 llama.py
```

## References
1. https://ai.facebook.com/blog/large-language-model-llama-meta-ai/
2. https://github.com/ggerganov/llama.cpp


