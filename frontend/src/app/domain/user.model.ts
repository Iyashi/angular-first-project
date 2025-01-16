export interface UserData {
  id: number;
  preName: string;
  lastName: string;
}

export class UserModel {
  public id: number = 0;
  public preName: string = '';
  public lastName: string = '';

  static fromData(data: UserData): UserModel {
    return new UserModel(data.id, data.preName, data.lastName);
  }

  constructor(id?: number, preName?: string, lastName?: string) {
    if (typeof id === 'number' && !isNaN(id)) this.id = id;
    if (typeof preName === 'string') this.preName = preName;
    if (typeof lastName === 'string') this.lastName = lastName;
  }

  toData(): UserData {
    return {
      id: this.id,
      preName: this.preName,
      lastName: this.lastName
    }
  }

  toString(): string {
    return `UserModel [id=${this.id}, preName=${this.preName}, lastName=${this.lastName}]`;
  }
}
