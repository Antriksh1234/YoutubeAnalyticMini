from transformers import pipeline

# Load the summarization pipeline with BART model
summarizer = pipeline("summarization", model="facebook/bart-large-cnn")

def lambda_handler(event, context):
    # Extract text from the event
    input_text = event.get('text', 'Default text if "text" not provided in event')

    # Generate summary using the loaded model
    generated_summary = summarize_text(input_text)

    return {
        'generated_summary': generated_summary
    }

def summarize_text(text):
    # Generate summary using the loaded model
    summary = summarizer(text, max_length=150, min_length=40, do_sample=False)[0]['summary_text']
    return summary
