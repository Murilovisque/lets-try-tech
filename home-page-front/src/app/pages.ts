import { Routes, Router } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { MembersComponent } from './members/members.component';

export class Pages {
    static readonly HOME_PATH = "home"
    static readonly MEMBERS_PATH = "members"

    static readonly ROUTES: Routes = [
        {path: "", redirectTo: Pages.HOME_PATH, pathMatch: "full"},
        {path: Pages.HOME_PATH, component: HomeComponent},
        {path: Pages.MEMBERS_PATH, component: MembersComponent}
    ];
    
    static navigateToMembersPage(router: Router): void {
        router.navigate([Pages.MEMBERS_PATH])
    }
}