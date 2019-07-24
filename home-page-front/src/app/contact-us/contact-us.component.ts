import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { CustomersService } from '../services/customers.service';
import { HttpBackend } from '@angular/common/http';

@Component({
	selector: 'app-contact-us',
	templateUrl: './contact-us.component.html',
	styleUrls: ['./contact-us.component.css']
})
export class ContactUsComponent implements OnInit {

	@Output()
	closePage = new EventEmitter<any>()
  	name :String
  	email :String
  	tel :Number
  	message :String

	constructor(private customersService: CustomersService) { }	

	ngOnInit() { }

	sendContactUsMessage(): void {
		this.customersService.sendContactUsMessage(this.name, this.tel, this.email, this.message).subscribe(
			() => console.log('ok'),
			(err) => console.log(err)
		)
	}
}
