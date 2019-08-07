import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { CustomersService } from '../services/customers.service';
import { BadRequestError, HomePageError } from '../errors/error';

@Component({
	selector: 'app-contact-us',
	templateUrl: './contact-us.component.html',
	styleUrls: ['./contact-us.component.css']
})
export class ContactUsComponent implements OnInit {

	@Output()
	closePage = new EventEmitter<any>()
  	name: string
  	email: string
  	tel: string
	message: string
	formMessage: FormMessage

	constructor(private customersService: CustomersService) { }	

	ngOnInit() {}

	sendContactUsMessage(): void {
		let telNum = this.tel != null ? Number(this.tel.replace(/[^0-9]/g, "")) : null;
		this.customersService.sendContactUsMessage(this.name, telNum, this.email, this.message).subscribe(
			() => this.formMessage = new FormMessage("Obrigado pelo seu contato. Clique para fechar", MessageType.OK),
			(err) => {
				if (err instanceof BadRequestError)
					this.formMessage = new FormMessage(err.message, MessageType.ERROR)
				else
					this.formMessage = new FormMessage((err as HomePageError).message + ". Clique para fechar", MessageType.FATAL)
			}
		)
	}

	clearFormMessage(element?: any): void {
		this.formMessage = null
	}

	closePopup(forceClose: boolean = false): void {
		if (forceClose)
			this.closePage.emit()
		else if (this.formMessage != null && (this.formMessage.type == MessageType.OK || this.formMessage.type == MessageType.FATAL))
			this.closePage.emit()
	}
}

class FormMessage {
	msg: string
	type: MessageType
	color: string

	constructor(msg: string, type: MessageType) {
		this.msg = msg
		this.type = type
		this.color = type == MessageType.OK ? "#1DADEE":  "#ec6048"
	}
}

enum MessageType { OK, ERROR, FATAL }