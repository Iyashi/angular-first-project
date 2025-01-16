import { Routes } from '@angular/router';

// Layouts
import { HomeLayout } from './layouts/home.layout';
import { DefaultLayout } from './layouts/default.layout';

// Pages
import { HomePage } from './pages/home.page';
import { UserPage } from './pages/user.page';
import { UsersPage } from './pages/users.page';

export const routes: Routes = [
  { // root page redirect to home page
    path: '',
    redirectTo: '/home',
    pathMatch: 'full',
  },

  { // home page using custom layout
    path: 'home',
    component: HomeLayout,
    children: [
      { path: '', component: HomePage }
    ]
  },

  { // other pages using default layout
    path: '',
    component: DefaultLayout,
    children: [
      { path: 'user/:user_id', component: UserPage },
      { path: 'users', component: UsersPage },
    ],
  },

  // { path: '**', component: PageNotFoundComponent }
];
