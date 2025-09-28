import { Product } from './product.model';

export interface ProductItem {
  id: number;
  productId: number;
  quantity: number;
  unitPrice: number;
  product?: Product;   // Transient : optionnel
}
