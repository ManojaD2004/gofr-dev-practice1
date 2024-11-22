function validateDataType(jsonData) {
    const allowedTypes = ["string", "int", "float32", "float64", "bool"];

    const validateValue = (value, key) => {
        if (typeof value === "string") {
            const seperateWords = value.split(" ");
            const firstWord = seperateWords[0];
            if (!allowedTypes.includes(firstWord)) {
                return { isValid: false, message: `Invalid data type for key: ${key}` };
            }
        } else if (Array.isArray(value)) {
            for (const item of value) {
                if (typeof item === "object" && item !== null) {
                    const nestedResult = validateDataType(item);
                    if (!nestedResult.isValid) {
                        return nestedResult;
                    }
                } else {
                    return { isValid: false, message: `Invalid value in array for key: ${key}` };
                }
            }
        } else if (typeof value === "object" && value !== null) {
            const nestedResult = validateDataType(value);
            if (!nestedResult.isValid) {
                return nestedResult;
            }
        }
        return { isValid: true };
    };
    for (const [key, value] of Object.entries(jsonData)) {
        const result = validateValue(value, key);
        if (!result.isValid) {
            return result;
        }
    }

    return { isValid: true, message: "All data types are valid." };
}

