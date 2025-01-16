import { Component } from '@angular/core';
import { NgIf } from '@angular/common';
import { UserService } from '../services/user.service'
import { UserModel } from '../domain/user.model';

@Component({
  imports: [NgIf],
  template: `
    <div>
      <h1>User List Data</h1>

      <ul *ngIf="users.length; else loading" class="user-list">
        @for (user of users; track user.id) {
          <li class="user-list-item">
            <pre>{{ user.id }}: {{ user.lastName }}, {{ user.preName }}</pre>
          </li>
        }
      </ul>
      <ng-template #loading>
        <p>Loading list of users...</p>
      </ng-template>
    </div>
  `,
})
export class UsersPage {
  users: UserModel[] = []

  constructor(private userService: UserService) {}

  ngOnInit() {
    this.userService.fetchUserList().subscribe(users => {
      this.users = users;
      console.log(users);
    });
  }
}
