import { Routes, Router } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { MembersComponent } from './members/members.component';
import { ContactUsComponent } from './contact-us/contact-us.component';

export class Pages {
    static readonly HOME_PATH = "home"
    static readonly MEMBERS_PATH = "membros"
    static readonly CONTACT_US_PATH = "contato"

    static readonly ROUTES: Routes = [
        {path: "", redirectTo: Pages.HOME_PATH, pathMatch: "full"},
        {path: Pages.HOME_PATH, component: HomeComponent},
        {path: Pages.MEMBERS_PATH, component: MembersComponent},
        {path: Pages.CONTACT_US_PATH, component: ContactUsComponent}
    ];
    
    static navigateToMembersPage(router: Router): void {
        router.navigate([Pages.MEMBERS_PATH])
    }

    static navigateToHomePage(router: Router): void {
        router.navigate([Pages.HOME_PATH])
    }

    static navigateToContactUs(router: Router): void {
        router.navigate([Pages.CONTACT_US_PATH])
    }
}