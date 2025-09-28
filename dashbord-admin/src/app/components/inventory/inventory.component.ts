import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Product } from '../../models/product.model';
import { InventoryService } from '../../services/inventory/inventory.service';

@Component({
  selector: 'app-inventory',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './inventory.component.html'
})
// ...existing code...
export class InventoryComponent implements OnInit {
  products: Product[] = [];
  editingProduct: Product | null = null;
  editForm = new FormGroup({
    name: new FormControl(''),
    quantity: new FormControl(0),
    price: new FormControl(0)
  });

  addingProduct = false;
  addForm = new FormGroup({
    name: new FormControl(''),
    quantity: new FormControl(0),
    price: new FormControl(0)
  });

  constructor(private inventoryService: InventoryService) {}

  ngOnInit(): void {
    this.loadProducts();
  }

  loadProducts(): void {
    this.inventoryService.getProducts().subscribe({
      next: (data) => this.products = data,
      error: (err) => console.error(err)
    });
  }

  // Ajout
  showAddForm(): void {
    this.addingProduct = true;
    this.addForm.reset();
    this.editingProduct = null;
  }

  saveNewProduct(): void {
   const newProduct: Product = {
     id: 0,
     name: this.addForm.value.name ?? '',
     quantity: Number(this.addForm.value.quantity) ?? 0,
     price: Number(this.addForm.value.price) ?? 0
    };
    this.inventoryService.addProduct(newProduct).subscribe({
      next: () => {
        this.addingProduct = false;
        this.loadProducts();
      },
      error: (err) => console.error(err)
    });
  }

  cancelAdd(): void {
    this.addingProduct = false;
  }

  // ...modification et suppression inchangées...
  editProduct(product: Product): void {
    this.editingProduct = product;
    this.editForm.setValue({
      name: product.name,
      quantity:product.quantity,
      price: product.price
    });
    this.addingProduct = false;
  }

  saveProduct(): void {
    if (this.editingProduct) {
      const updatedProduct: Product = {
      id: this.editingProduct.id,
      name: this.editForm.value.name ?? '',
      quantity: Number(this.editForm.value.quantity) ?? 0,
      price: Number(this.editForm.value.price) ?? 0
    };
      this.inventoryService.updateProduct(updatedProduct.id, updatedProduct).subscribe({
        next: () => {
          this.editingProduct = null;
          this.loadProducts();
        },
        error: (err) => console.error(err)
      });
    }
  }

  cancelEdit(): void {
    this.editingProduct = null;
  }

  deleteProduct(id: number): void {
    if (confirm('Voulez-vous vraiment supprimer ce client ?')) {
      this.inventoryService.deleteProduct(id).subscribe({
        next: () => this.loadProducts(),
        error: (err) => console.error(err)
      });
    }
  }
}
// ...existing code...