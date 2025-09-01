let amisLib = amisRequire('amis');
let React = amisRequire('react');

function SSEComponent(props) {
    const containerRef = React.useRef(null);
    const [state, setState] = React.useState({
        data: null,
        loading: true,
        error: null
    });
    const eventSourceRef = React.useRef(null);
    // 仅用于标记是否首次渲染（核心：确保首次渲染执行）
    const isFirstRender = React.useRef(true);

    // 建立SSE连接
    const connect = React.useCallback(() => {
        if (eventSourceRef.current) {
            eventSourceRef.current.close();
        }

        const { url } = props;
        if (!url) {
            setState(prev => ({ ...prev, error: '缺少SSE服务地址', loading: false }));
            return;
        }

        // 重连时仅更新loading，保留数据
        setState(prev => ({ ...prev, loading: true, error: null }));

        try {
            const eventSource = new EventSource(url);
            eventSourceRef.current = eventSource;

            eventSource.onopen = () => {
                setState(prev => ({ ...prev, loading: false }));
            };

            eventSource.onmessage = (event) => {
                try {
                    const data = JSON.parse(event.data);
                    setState(prev => ({
                        ...prev,
                        data,
                        loading: false,
                        error: null
                    }));
                } catch (e) {
                    setState(prev => ({
                        ...prev,
                        error: '数据解析失败',
                        loading: false
                    }));
                    console.error('SSE数据解析错误:', e);
                }
            };

            eventSource.onerror = () => {
                setState(prev => ({
                    ...prev,
                    error: '连接异常，正在重试...',
                    loading: true
                }));
                if (eventSourceRef.current) {
                    eventSourceRef.current.close();
                    setTimeout(connect, 3000);
                }
            };
        } catch (e) {
            setState(prev => ({
                ...prev,
                error: '初始化连接失败',
                loading: false
            }));
            console.error('SSE连接建立失败:', e);
        }
    }, [props.url]);

    // 组件挂载时立即初始化（确保首次渲染触发）
    React.useEffect(() => {
        connect();
        return () => {
            if (eventSourceRef.current) {
                eventSourceRef.current.close();
            }
        };
    }, [connect]);

    // 渲染内容：完全保留原始可渲染逻辑，仅添加重连优化
    const renderContent = () => {
        const { loading, error, data } = state;
        const { loading: loadingTpl, error: errorTpl, body } = props;

        const scope = {
            ...data,
            $loading: loading,
            $error: error
        };

        const context = {
            ...(amisLib.context || {}),
            data: scope,
            props
        };

        // 错误状态（修复模板误用）
        if (error) {
            return errorTpl
                ? amisLib.render(errorTpl, context)
                : React.createElement('div', { style: { color: 'red' } }, `错误: ${error}`);
        }

        // 加载状态
        if (loading) {
            return loadingTpl
                ? amisLib.render(loadingTpl, context)
                : React.createElement('div', null, '加载中...');
        }

        // 主体内容：首次渲染直接执行，后续复用结果
        if (body) {
            // 首次渲染：直接返回amis渲染结果（确保显示）
            if (isFirstRender.current) {
                const firstRenderResult = amisLib.render(body, context);
                isFirstRender.current = false; // 标记首次渲染完成
                return firstRenderResult;
            } else {
                // 后续更新：用新数据重新渲染到同一结构（不重建组件）
                return amisLib.render(body, context);
            }
        }

        return React.createElement('div', null, JSON.stringify(data, null, 2));
    };

    // 直接返回渲染结果（原始可渲染的核心逻辑）
    return React.createElement('div', { ref: containerRef }, renderContent());
}

// 注册自定义组件
amisLib.Renderer({
    test: /(^|\/)sse$/
})(SSEComponent);