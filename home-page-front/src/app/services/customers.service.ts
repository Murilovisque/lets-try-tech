import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, of, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Errors } from '../errors/error';

@Injectable({
  providedIn: 'root'
})
export class CustomersService {

  	constructor(private http: HttpClient) { }

	sendContactUsMessage(name: string, tel: number, email: string, message: string): Observable<any> {
		const url = "api/customers/contact-message"
		const httpOptions = { headers: new HttpHeaders({ 'Content-Type': 'application/json' }) };
		let body = {
			name: name,
			tel:  tel,
			email: email,
			message: message
		}
    	return this.http.post<any>(url, JSON.stringify(body), httpOptions).pipe(catchError((err) => Errors.getObservableError(err)))
  	}
}
