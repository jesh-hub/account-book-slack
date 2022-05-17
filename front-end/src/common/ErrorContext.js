import React from 'react';

const errors = [];

function addError(err) {
  const nextId = errors.length > 0 ? errors[errors.length - 1].id + 1 : 0;
  errors.push({
    id: nextId,
    code: err.code,
    message: err.message
  });
}

function deleteError(err) {
  const index = errors.indexOf(err);
  if (index >= 0)
    errors.splice(index, 1);
}

export const ErrorContext = React.createContext({
  errors,
  addError,
  deleteError
});

export default ErrorContext;
