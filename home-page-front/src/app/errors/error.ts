import { HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';

export class Errors {
    static getObservableError(err :any) :Observable<never> {
        if (err instanceof HttpErrorResponse && err.status == 400)
            return throwError(new BadRequestError(err.error))
        return throwError(new HomePageError())
    }
}

export class HomePageError implements Error {
    name: string;
    message: string;
    stack?: string;

    constructor(message? :string) {
        this.message = message == null ? "Não foi possível processar sua requisição, tente novamente mais tarde" : message;
        this.name = this.getName();
    }

    getName() :string {
        return "HomePageError";
    }
}

export class BadRequestError extends HomePageError {
    constructor(message :string) {
        super(message)
    }

    getName() :string {
        return "BadRequestError";
    }
}