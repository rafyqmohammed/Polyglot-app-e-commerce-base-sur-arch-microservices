package com.example.demo.billingservice.repository;

import com.example.demo.billingservice.entities.Bill;
import org.springframework.data.jpa.repository.JpaRepository;

public interface BillRepository  extends JpaRepository<Bill, Long> {
}