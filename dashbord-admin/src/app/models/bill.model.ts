import { ProductItem } from './product-item.model';
import { Customer } from './customer.model';

export interface Bill {
  id: number;
  billingDate: Date;        // Java Date → TypeScript Date
  customerId: number;
  productItems: ProductItem[];
  customer?: Customer;      // Transient : optionnel, reçu si exposé dans l’API
}
