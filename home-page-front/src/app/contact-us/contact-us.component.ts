import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { CustomersService } from '../services/customers.service';
import { HttpBackend } from '@angular/common/http';
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
  	tel: number
	message: string
	formMessage: FormMessage

	constructor(private customersService: CustomersService) { }	

	ngOnInit() { }

	sendContactUsMessage(): void {
		this.customersService.sendContactUsMessage(this.name, this.tel, this.email, this.message).subscribe(
			() => this.formMessage = new FormMessage("Obrigado pelo seu contato. Clique para fechar", MessageType.OK),
			(err) => {
				if (err instanceof BadRequestError)
					this.formMessage = new FormMessage(err.message, MessageType.ERROR)
				else
					this.formMessage = new FormMessage((err as HomePageError).message, MessageType.FATAL)
			}
		)
	}

	clearFormMessage(): void {
		this.formMessage = null
	}

	closePopup(forceClose: boolean = false): void {
		if (forceClose)
			this.closePage.emit()
		else if (this.formMessage != null && (this.formMessage.type == MessageType.OK || this.formMessage.type == MessageType.FATAL))
			this.closePage.emit()
	}

	setTel(newTel: any): void {
		if (/^\d+$/.test(newTel.value))
			this.tel = Number(newTel.value)
		else
			newTel.value = this.tel == null ? "":  this.tel
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