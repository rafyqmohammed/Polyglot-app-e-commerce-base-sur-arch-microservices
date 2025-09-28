package com.example.demo.billingservice.web;

import com.example.demo.billingservice.entities.Bill;
import com.example.demo.billingservice.feign.CustomerRestClient;
import com.example.demo.billingservice.feign.ProductRestClient;
import com.example.demo.billingservice.model.Customer;
import com.example.demo.billingservice.repository.BillRepository;
import com.example.demo.billingservice.repository.ProductItemRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
@RestController
public class BillRestController {
    @Autowired
    private BillRepository billRepository;
    @Autowired
    private ProductItemRepository productItemRepository;
    @Autowired
    private CustomerRestClient customerRestClient;
    @Autowired
    private ProductRestClient productRestClient;
    @GetMapping(path = "/bills/{id}")
    public Bill getBill(@PathVariable Long id){
        Bill bill = billRepository.findById(id).get();
        bill.setCustomer(customerRestClient.getCustomerById(bill.getCustomerId()));
        bill.getProductItems().forEach(productItem -> {
            productItem.setProduct(productRestClient.getProductById(productItem.getProductId()));
        });
        return bill;
    }
    @GetMapping( "/bills")
    public List<Bill> getBills(){
        List<Bill> bills = billRepository.findAll();
        bills.forEach(bill -> {
            bill.setCustomer(customerRestClient.getCustomerById(bill.getCustomerId()));
            bill.getProductItems().forEach(productItem -> {
                productItem.setProduct(productRestClient.getProductById(productItem.getProductId()));
            });

        });
        return bills;
    }
    @GetMapping( "/billsall")
    public List<Bill> getBillsall(){
        List<Bill> bills = billRepository.findAll();

        return bills;
    }

}