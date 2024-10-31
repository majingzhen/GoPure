/**
 * HTTP请求工具类
 */
(function($) {
    const request = {
        /**
         * GET请求
         * @param {string} url - 请求地址
         * @param {object} params - 请求参数
         * @param {object} options - 其他选项
         * @returns {Promise}
         */
        get: function(url, params = {}, options = {}) {
            // 如果有参数，将参数拼接到URL后面
            const queryString = this.buildQueryString(params);
            const finalUrl = queryString ? `${url}${url.includes('?') ? '&' : '?'}${queryString}` : url;

            return this.request({
                url: finalUrl,
                method: 'GET',
                ...options
            });
        },

        /**
         * 构建查询字符串
         * @param {object} params - 参数对象
         * @returns {string} 查询字符串
         */
        buildQueryString: function(params) {
            if (!params || Object.keys(params).length === 0) {
                return '';
            }
            return Object.entries(params)
                .filter(([_, value]) => value != null && value !== '') // 过滤掉null、undefined和空字符串
                .map(([key, value]) => {
                    if (Array.isArray(value)) {
                        // 处理数组参数
                        return value.map(v => `${encodeURIComponent(key)}=${encodeURIComponent(v)}`).join('&');
                    }
                    return `${encodeURIComponent(key)}=${encodeURIComponent(value)}`;
                })
                .join('&');
        },

        /**
         * POST请求 - JSON格式
         * @param {string} url - 请求地址
         * @param {object} data - 请求数据
         * @param {object} options - 其他选项
         * @returns {Promise}
         */
        post: function(url, data = {}, options = {}) {
            return this.request({
                url: url,
                method: 'POST',
                data: data,
                contentType: 'application/json',
                ...options
            });
        },

        /**
         * POST请求 - 表单格式
         * @param {string} url - 请求地址
         * @param {object} data - 请求数据
         * @param {object} options - 其他选项
         * @returns {Promise}
         */
        form: function(url, data = {}, options = {}) {
            return this.request({
                url: url,
                method: 'POST',
                data: data,
                contentType: 'application/x-www-form-urlencoded',
                processData: true,
                ...options
            });
        },

        /**
         * 统一请求方法
         * @param {object} config - 请求配置
         * @returns {Promise}
         */
        request: function(config) {
            const defaultConfig = {
                dataType: 'json',
                timeout: 10000,
                contentType: 'application/json',
                processData: false
            };

            // 合并配置
            const finalConfig = {
                ...defaultConfig,
                ...config
            };

            // 处理请求数据
            if (finalConfig.contentType === 'application/json' && typeof finalConfig.data === 'object') {
                finalConfig.data = JSON.stringify(finalConfig.data);
            }

            return new Promise((resolve, reject) => {
                $.ajax({
                    ...finalConfig,
                    success: (response) => {
                        if (response.code === 0) {
                            resolve(response);
                        } else {
                            this.handleError(response);
                            reject(response);
                        }
                    },
                    error: (xhr, textStatus, error) => {
                        if (xhr.status === 401) {
                            // 未授权，重定向到登录页
                            this.redirectToLogin();
                        } else {
                            const errorMsg = this.getErrorMessage(xhr, textStatus, error);
                            this.handleError({ msg: errorMsg });
                            reject({ code: xhr.status, msg: errorMsg });
                        }
                    }
                });
            });
        },

        /**
         * 处理错误
         * @param {object} response - 响应对象
         */
        handleError: function(response) {
            if (response.msg) {
                // 使用layui的提示
                if (layui && layui.layer) {
                    layui.layer.msg(response.msg, { icon: 2 });
                } else {
                    alert(response.msg);
                }
            }
        },

        /**
         * 获取错误信息
         * @param {object} xhr - XMLHttpRequest对象
         * @param {string} textStatus - 状态文本
         * @param {string} error - 错误信息
         * @returns {string}
         */
        getErrorMessage: function(xhr, textStatus, error) {
            if (xhr.responseJSON && xhr.responseJSON.msg) {
                return xhr.responseJSON.msg;
            }
            switch (xhr.status) {
                case 400:
                    return '请求参数错误';
                case 401:
                    return '未授权，请重新登录';
                case 403:
                    return '拒绝访问';
                case 404:
                    return '请求地址不存在';
                case 500:
                    return '服务器内部错误';
                default:
                    return '网络错误，请稍后重试';
            }
        },

        /**
         * 重定向到登录页
         */
        redirectToLogin: function() {
            // 如果已经在登录页，不需要重定向
            if (window.location.pathname === '/login') {
                return;
            }
            // 保存当前页面URL，登录后可以跳回来
            const currentPath = encodeURIComponent(window.location.pathname + window.location.search);
            window.location.href = `/login?redirect=${currentPath}`;
        }
    };

// 添加全局Ajax设置
    $.ajaxSetup({
        complete: function(xhr) {
            if (xhr.status === 401) {
                request.redirectToLogin();
            }
        }
    });

// 将request对象挂载到window上，使其全局可用
    window.request = request;
})(window.jQuery);
