import { Component } from '@angular/core';
import { NgIf } from '@angular/common';
import { Router, ActivatedRoute } from '@angular/router';
import { UserService } from '../services/user.service'
import { UserModel } from '../domain/user.model';

@Component({
  imports: [NgIf],
  template: `
    <div>
      <h1>User Data</h1>

      <div *ngIf="user !== null; else loading" class="user">
        <pre>{{ user.id }}: {{ user.lastName }}, {{ user.preName }}</pre>
      </div>
      <ng-template #loading>
        <p>Loading user...</p>
      </ng-template>
    </div>
  `,
})
export class UserPage {
  user: UserModel | null = null;

  constructor(private route: ActivatedRoute, private userService: UserService) {
    this.route.params.subscribe(params => {
      const id = parseInt(params['user_id'] || '', 10) || 0
      console.log('Route :user_id changed', id);
      this.userService.fetchUser(id).subscribe(user => {
        this.user = user
        console.log(user)
      })
    });
  }
}
