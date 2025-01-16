import { Routes } from '@angular/router';

// Layouts
import { HomeLayout } from './layouts/home.layout';
import { DefaultLayout } from './layouts/default.layout';

// Pages
import { HomePage } from './pages/home.page';

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
  },

  // { path: '**', component: PageNotFoundComponent }
];
