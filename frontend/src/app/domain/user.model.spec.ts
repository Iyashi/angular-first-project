import { UserModel } from './user.model';

describe('UserModel', () => {
  const user = new UserModel(1, 'John', 'Doe')

  it('should create an instance', () => {
    expect(user).toBeTruthy();
  });

  it('should set id via constructor', () => {
    expect(user.id).toBe(1);
  });
  it('should set preName via constructor', () => {
    expect(user.preName).toBe('John');
  });
  it('should set lastName via constructor', () => {
    expect(user.preName).toBe('Doe');
  });

  it('should convert from data correctly', () => {
    expect(UserModel.fromData({ id: 1, preName: 'John', lastName: 'Doe' })).toEqual(user);
  });
  it('should convert to data correctly', () => {
    expect(user.toData()).toBe({ id: 1, preName: 'John', lastName: 'Doe' });
  });

  it('should format as string correctly', () => {
    expect(user.toString()).toBe('UserModel [id=1, preName=John, lastName=Doe]');
  });
});
