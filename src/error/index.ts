type OperationalErrorOptions = {
  code?: string;
  [key: string]: unknown;
};

interface OperationalError extends Error {
  name: string;
  isOperational: boolean;
  details?: string;
  [key: string]: unknown;
}

export function g3Error(
  message: string,
  details: string = "",
  options: OperationalErrorOptions = {},
): OperationalError {
  const error = new Error(message) as OperationalError;
  error.name = "G3Error";
  error.isOperational = true;
  error.details = details;
  Object.assign(error, options);
  Error.captureStackTrace(error, g3Error);
  return error;
}
