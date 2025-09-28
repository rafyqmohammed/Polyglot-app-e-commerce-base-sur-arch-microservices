package com.example.demo.billingservice.entities;

import jakarta.persistence.*;
import lombok.*;
import com.example.demo.billingservice.model.Customer;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;
@Entity
@NoArgsConstructor @AllArgsConstructor @Getter @Setter @Builder
public class Bill {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private Date billingDate;
    private int customerId;
    @OneToMany(mappedBy = "bill")
    private List<ProductItem> productItems = new ArrayList<>();
    //
    @Transient private Customer customer;

    public Long getId() {
        return id;
    }
    public void setId(Long id) {
        this.id = id;
    }
    public Date getBillingDate() {
        return billingDate;
    }
    public void setBillingDate(Date billingDate) {
        this.billingDate = billingDate;
    }
    public int getCustomerId() {
        return customerId;
    }
    public void setCustomerId(int customerId) {
        this.customerId = customerId;
    }
    public List<ProductItem> getProductItems() {
        return productItems;
    }
    public void setProductItems(List<ProductItem> productItems) {
        this.productItems = productItems;
    }
    public Customer getCustomer() {
        return customer;
    }
    public void setCustomer(Customer customer) {
        this.customer = customer;
    }

}