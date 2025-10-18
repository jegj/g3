type OperationalErrorOptions = {
  code?: string;
  [key: string]: unknown;
};

export interface OperationalError extends Error {
  name: string;
  isOperational: boolean;
  [key: string]: unknown;
}

export function g3Error(
  message: string,
  options: OperationalErrorOptions = {},
): OperationalError {
  const error = new Error(message) as OperationalError;
  error.name = "G3Error";
  error.isOperational = true;
  Object.assign(error, options);
  Error.captureStackTrace(error, g3Error);
  return error;
}
