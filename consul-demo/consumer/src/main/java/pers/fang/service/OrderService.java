package pers.fang.service;

import pers.fang.entity.Order;

/**
 * @author FANG
 * @create 2022-03-26 16:33
 */
public interface OrderService {
    Order selectOrderById(Integer id);
}
