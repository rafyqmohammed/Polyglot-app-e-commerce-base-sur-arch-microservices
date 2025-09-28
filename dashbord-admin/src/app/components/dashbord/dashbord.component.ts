import { Component } from '@angular/core';
import {RouterLink} from "@angular/router";
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [
    RouterLink,
  RouterOutlet
  ],
  templateUrl: './dashbord.component.html' // ← le fichier HTML que je vous ai donné
})
export class DashbordComponent {
  constructor() {}

  // logout() {
  //   // Exemple simple : supprimer le token de l'utilisateur
  //   localStorage.removeItem('auth_token');

  //   // Rediriger vers la page de login
  //   this.router.navigate(['/login']);
  // }
}
