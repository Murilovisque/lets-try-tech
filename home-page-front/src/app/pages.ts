import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';

export class Pages {
    static readonly HOME_PATH = "home";

    static readonly ROUTES: Routes = [
        {path: "", redirectTo: Pages.HOME_PATH, pathMatch: "full"},
        {path:Pages.HOME_PATH, component: HomeComponent}
    ];
    

}