import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UserModel } from '../domain/user.model';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  static get ApiUrl() { return new URL('http://api.angular-first-project.local') };

  constructor(private http: HttpClient) {}

  fetchUser(id: number) {
    const url = UserService.ApiUrl
    url.pathname += `user/${id}`;
    const response = this.http.get<UserModel>(url.toString());
    return response;
  }

  fetchUserList() {
    const url = UserService.ApiUrl
    url.pathname += `users`;
    const response = this.http.get<UserModel[]>(url.toString());
    return response;
  }
}
