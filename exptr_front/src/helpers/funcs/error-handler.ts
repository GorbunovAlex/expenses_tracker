import { Notify } from 'quasar';

class ErrorHandler {
  error: Error;

  constructor(error: Error) {
    this.error = error;
  }

  handleError() {
    Notify.create({
      message: this.error.message,
      color: 'negative',
      icon: 'report_problem',
      position: 'top-right',
    });
  }
}

export default ErrorHandler;
