/**
 * Formats error messages by capitalizing the first letter and adding a period if missing
 * @param message - The error message to format
 * @returns The formatted error message
 */
export function formatErrorMessage(message: string | null | undefined): string {
    if (!message || typeof message !== 'string') {
        return "An error occurred.";
    }
    
    // Trim whitespace
    let formattedMessage = message.trim();
    
    // Capitalize first letter
    if (formattedMessage.length > 0) {
        formattedMessage = formattedMessage.charAt(0).toUpperCase() + formattedMessage.slice(1);
    }
    
    // Add period if it doesn't end with punctuation
    if (formattedMessage.length > 0 && !formattedMessage.match(/[.!?]$/)) {
        formattedMessage += '.';
    }
    
    return formattedMessage;
}
