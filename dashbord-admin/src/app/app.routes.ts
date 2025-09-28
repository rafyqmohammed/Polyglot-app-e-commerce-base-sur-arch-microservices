import { Routes } from '@angular/router';
import { CustomerComponent } from './components/customer/customer.component';
import { BillingComponent } from './components/billing/billing.component';
import { InventoryComponent } from './components/inventory/inventory.component';
import { DashbordComponent } from './components/dashbord/dashbord.component';

export const routes: Routes = [
  {
    path: '',
    component: DashbordComponent,
    children: [
      { path: 'customers', component: CustomerComponent },
      { path: 'billing', component: BillingComponent },
      { path: 'inventory', component: InventoryComponent }
    ]
  }
];
