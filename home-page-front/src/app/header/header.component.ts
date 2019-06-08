import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Pages } from '../pages';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  showContactUsPage = false

  constructor(
    private router: Router
  ) { }

  ngOnInit() {}

  linkMembersPage(): void {
    Pages.navigateToMembersPage(this.router)
  }

  linkHomePage(): void {
    Pages.navigateToHomePage(this.router)
  }

  linkContactUs(): void {
    this.showContactUsPage = true
  }

  onCloseContactUsPage() : void {
    this.showContactUsPage = false
  }
}
