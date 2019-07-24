import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class CustomersService {

  	constructor(private http: HttpClient) { }

	sendContactUsMessage(name :String, tel :Number, email :String, message :String): Observable<any> {
		const url = "api/customers/contact-message"
		const httpOptions = { headers: new HttpHeaders({ 'Content-Type': 'application/json' }) };
		let body = {
			name: name,
			tel:  tel,
			email: email,
			message: message
		}
    	return this.http.post<any>(url, body, httpOptions).pipe(catchError(err => throwError(err)))
  	}
}
