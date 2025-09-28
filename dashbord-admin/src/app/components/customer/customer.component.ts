import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { CustomerService } from '../../services/customer/customer.service';
import { Customer } from '../../models/customer.model';

@Component({
  selector: 'app-customer',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './customer.component.html'
})
// ...existing code...
export class CustomerComponent implements OnInit {
  customers: Customer[] = [];
  editingCustomer: Customer | null = null;
  editForm = new FormGroup({
    name: new FormControl(''),
    email: new FormControl('')
  });

  addingCustomer = false;
  addForm = new FormGroup({
    name: new FormControl(''),
    email: new FormControl('')
  });

  constructor(private customerService: CustomerService) {}

  ngOnInit(): void {
    this.loadCustomers();
  }

  loadCustomers(): void {
    this.customerService.getCustomers().subscribe({
      next: (data) => this.customers = data,
      error: (err) => console.error(err)
    });
  }

  // Ajout
  showAddForm(): void {
    this.addingCustomer = true;
    this.addForm.reset();
    this.editingCustomer = null;
  }

  saveNewCustomer(): void {
    const newCustomer: Customer = {
      id: 0, // ou undefined si géré par le backend
      name: this.addForm.value.name ?? '',
      email: this.addForm.value.email ?? ''
    };
    this.customerService.addCustomer(newCustomer).subscribe({
      next: () => {
        this.addingCustomer = false;
        this.loadCustomers();
      },
      error: (err) => console.error(err)
    });
  }

  cancelAdd(): void {
    this.addingCustomer = false;
  }

  // ...modification et suppression inchangées...
  editCustomer(customer: Customer): void {
    this.editingCustomer = customer;
    this.editForm.setValue({
      name: customer.name,
      email: customer.email
    });
    this.addingCustomer = false;
  }

  saveCustomer(): void {
    if (this.editingCustomer) {
      const updatedCustomer: Customer = {
        id: this.editingCustomer.id,
        name: this.editForm.value.name ?? '',
        email: this.editForm.value.email ?? ''
      };
      this.customerService.updateCustomer(updatedCustomer.id, updatedCustomer).subscribe({
        next: () => {
          this.editingCustomer = null;
          this.loadCustomers();
        },
        error: (err) => console.error(err)
      });
    }
  }

  cancelEdit(): void {
    this.editingCustomer = null;
  }

  deleteCustomer(id: number): void {
    if (confirm('Voulez-vous vraiment supprimer ce client ?')) {
      this.customerService.deleteCustomer(id).subscribe({
        next: () => this.loadCustomers(),
        error: (err) => console.error(err)
      });
    }
  }
}
// ...existing code...