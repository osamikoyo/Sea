#include <stdio.h>
#include <string.h>
#include <ctype.h>

int valid(char* text) {
	if (text == NULL) {
		return 0;
	}

	int len = strlen(text);
	int inString = 0;
	int inArray = 0;
	int inTable = 0;
	int lineStart = 1;
	
	for (int i = 0; i < len; i++) {
		char c = text[i];
		
		// Handle strings
		if (c == '"' && (i == 0 || text[i-1] != '\\')) {
			inString = !inString;
			continue;
		}
		
		// Skip processing if we're inside a string
		if (inString) {
			continue;
		}
		
		// Handle comments
		if (c == '#' && !inString) {
			while (i < len && text[i] != '\n') {
				i++;
			}
			lineStart = 1;
			continue;
		}
		
		// Handle arrays
		if (c == '[') {
			if (i < len - 1 && text[i+1] == '[') {
				// Double bracket for table array
				if (inTable || inArray) return 0;
				i++;
				inTable = 1;
			} else {
				inArray++;
			}
		}
		if (c == ']') {
			if (i < len - 1 && text[i+1] == ']') {
				// Double bracket for table array
				if (!inTable) return 0;
				i++;
				inTable = 0;
			} else {
				if (inArray > 0) inArray--;
				else return 0;
			}
		}
		
		// Check for invalid characters at start of line
		if (lineStart && !isspace(c) && c != '[' && c != '#') {
			// Key must start with letter, number or underscore
			if (!isalnum(c) && c != '_') {
				return 0;
			}
			lineStart = 0;
		}
		
		// Handle newlines
		if (c == '\n') {
			lineStart = 1;
		}
		
		// Basic key-value separator check
		if (c == '=' && !inArray && !inTable) {
			if (i > 0 && isspace(text[i-1]) && i < len-1 && isspace(text[i+1])) {
				continue;
			}
			// Must have space before and after =
			return 0;
		}
	}
	
	// Check for unclosed constructs
	if (inString || inArray || inTable) {
		return 0;
	}
	
	return 1;
}
