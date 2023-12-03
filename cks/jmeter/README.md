### 启动服务

```
sudo /usr/local/apache-jmeter-5.6.2/bin/jmeter -v
sudo /usr/local/apache-jmeter-5.6.2/bin/jmeter-server &
```

### 参考文件

```
https://takuto-n.github.io/post/20200614-jmeter-master-slave/
https://qiita.com/miyuki_samitani/items/78080ec4a2fa1e150241
```





`https://hello-bi7moqtv2q-uc.a.run.app/`

**停用 AppArmor 服务：**

```
sudo systemctl stop apparmor
```

**禁用 AppArmor 服务的自动启动：**

```
sudo systemctl disable apparmor
sudo systemctl status apparmor
```

**重启系统：**

```
sudo reboot
```

检查 `ufw` 状态：

```
sudo ufw status
```

确保 `ufw` 已安装并处于活动状态。

**如果已启用，禁用 `ufw`：**

```
sudo ufw disable
```

#### 对于 ufw：

```
sudo ufw disable
```

#### 对于 firewalld：

```
sudo systemctl stop firewalld
sudo systemctl disable firewalld
```

你可以使用以下命令来检查系统上是否安装了 `iptables`：

```
sudo iptables --version
```

查看当前的 `iptables` 规则：

```
sudo iptables -L
```

这会列出当前的 iptables 规则。

**清除所有规则：**

```
sudo iptables -F
```

这会清除所有防火墙规则。

**允许所有流量通过：**

```
sudo iptables -P INPUT ACCEPT
sudo iptables -P FORWARD ACCEPT
sudo iptables -P OUTPUT ACCEPT
```



配置机器主机名
 在 192.168.40.180 上执行如下:
 hostnamectl set-hostname xianchaomaster1 && bash 在 192.168.40.181 上执行如下:
 hostnamectl set-hostname xianchaonode1 && bash 在 192.168.40.182 上执行如下:
 hostnamectl set-hostname xianchaonode2 && bash



配置主机 hosts 文件，相互之间通过主机名互相访问 修改每台机器的/etc/hosts 文件，增加如下三行: 192.168.40.180 xianchaomaster1 192.168.40.181 xianchaonode1

192.168.40.182 xianchaonode2

[root@xianchaomaster1 ~]# ssh-keygen #一路回车，不输入密码 把本地生成的密钥文件和私钥文件拷贝到远程主机 [root@xianchaomaster1 ~]# ssh-copy-id xianchaomaster1





```
andre@jmst:~$ cd /usr/local/
andre@jmst:/usr/local$ sudo wget https://dlcdn.apache.org//jmeter/binaries/apache-jmeter-5.6.2.zip
vim /usr/local/apache-jmeter-5.6.2/bin/jmeter.properties
sudo vim  /etc/hosts
sudo /usr/local/apache-jmeter-5.6.2/bin/jmeter -v

sudo /usr/local/apache-jmeter-5.6.2/bin/jmeter-server &

ps ax | grep jmeter
```





```
sudo vim /usr/local/apache-jmeter-5.6.2/bin/jmeter-server
sudo vim /usr/local/apache-jmeter-5.6.2/bin/jmeter.properties
sudo vim /etc/hosts
sudo /usr/local/apache-jmeter-5.6.2/bin/jmeter -v
sudo /usr/local/apache-jmeter-5.6.2/bin/jmeter-server &
```



```
<?xml version="1.0" encoding="UTF-8"?>
<jmeterTestPlan version="1.2" properties="5.0" jmeter="5.4.1">
  <hashTree>
    <TestPlan guiclass="TestPlanGui" testclass="TestPlan" testname="Test Plan" enabled="true">
      <stringProp name="TestPlan.comments"></stringProp>
      <boolProp name="TestPlan.functional_mode">false</boolProp>
      <boolProp name="TestPlan.tearDown_on_shutdown">true</boolProp>
      <boolProp name="TestPlan.serialize_threadgroups">false</boolProp>
      <elementProp name="TestPlan.user_defined_variables" elementType="Arguments" guiclass="ArgumentsPanel" testclass="Argume
nts" testname="User Defined Variables" enabled="true">
        <collectionProp name="Arguments.arguments"/>
      </elementProp>
      <stringProp name="TestPlan.user_define_classpath"></stringProp>
    </TestPlan>
    <hashTree>
      <Arguments guiclass="ArgumentsPanel" testclass="Arguments" testname="User Defined Variables" enabled="true">
        <collectionProp name="Arguments.arguments">
          <elementProp name="url" elementType="Argument">
            <stringProp name="Argument.name">url</stringProp>
            <stringProp name="Argument.value">[=url]</stringProp>
            <stringProp name="Argument.metadata">=</stringProp>
          </elementProp>
          <elementProp name="contentType" elementType="Argument">
            <stringProp name="Argument.name">contentType</stringProp>
            <stringProp name="Argument.value">[=contentType]</stringProp>
            <stringProp name="Argument.metadata">=</stringProp>
          </elementProp>
          <elementProp name="method" elementType="Argument">
            <stringProp name="Argument.name">method</stringProp>
            <stringProp name="Argument.value">[=method]</stringProp>
            <stringProp name="Argument.metadata">=</stringProp>
          </elementProp>
          <elementProp name="body" elementType="Argument">
            <stringProp name="Argument.name">body</stringProp>
            <stringProp name="Argument.value">[=body]</stringProp>
            <stringProp name="Argument.metadata">=</stringProp>
          </elementProp>
        </collectionProp>
      </Arguments>
      <hashTree/>
      <ResultCollector guiclass="ViewResultsFullVisualizer" testclass="ResultCollector" testname="View Results Tree" enabled=
"true">
        <boolProp name="ResultCollector.error_logging">false</boolProp>
        <objProp>
          <name>saveConfig</name>
          <value class="SampleSaveConfiguration">
            <time>true</time>
            <latency>true</latency>
            <timestamp>true</timestamp>
            <success>true</success>
            <label>true</label>
            <code>true</code>
            <message>true</message>
            <threadName>true</threadName>
            <dataType>true</dataType>
            <encoding>false</encoding>
            <assertions>true</assertions>
            <subresults>true</subresults>
            <responseData>false</responseData>
            <samplerData>false</samplerData>
            <xml>true</xml>
            <fieldNames>true</fieldNames>
            <responseHeaders>false</responseHeaders>
            <requestHeaders>false</requestHeaders>
            <responseDataOnError>false</responseDataOnError>
            <saveAssertionResultsFailureMessage>true</saveAssertionResultsFailureMessage>
            <assertionsResultsToSave>0</assertionsResultsToSave>
            <bytes>true</bytes>
            <sentBytes>true</sentBytes>
            <url>true</url>
            <threadCounts>true</threadCounts>
            <idleTime>true</idleTime>
            <connectTime>true</connectTime>
          </value>
        </objProp>
        <stringProp name="filename"></stringProp>
      </ResultCollector>
      <hashTree/>
      <ThreadGroup guiclass="ThreadGroupGui" testclass="ThreadGroup" testname="Thread Group" enabled="true">
        <stringProp name="ThreadGroup.on_sample_error">continue</stringProp>
        <elementProp name="ThreadGroup.main_controller" elementType="LoopController" guiclass="LoopControlPanel" testclass="L
oopController" testname="Loop Controller" enabled="true">
          <boolProp name="LoopController.continue_forever">false</boolProp>
          <stringProp name="LoopController.loops">1</stringProp>
        </elementProp>
        <stringProp name="ThreadGroup.num_threads">1</stringProp>
        <stringProp name="ThreadGroup.ramp_time">1</stringProp>
        <boolProp name="ThreadGroup.scheduler">false</boolProp>
        <stringProp name="ThreadGroup.duration"></stringProp>
        <stringProp name="ThreadGroup.delay"></stringProp>
        <boolProp name="ThreadGroup.same_user_on_next_iteration">true</boolProp>
      </ThreadGroup>
      <hashTree>
        <HTTPSamplerProxy guiclass="HttpTestSampleGui" testclass="HTTPSamplerProxy" testname="HTTP Request" enabled="true">
          <boolProp name="HTTPSampler.postBodyRaw">true</boolProp>
          <elementProp name="HTTPsampler.Arguments" elementType="Arguments">
            <collectionProp name="Arguments.arguments">
              <elementProp name="" elementType="HTTPArgument">
                <boolProp name="HTTPArgument.always_encode">false</boolProp>
                <stringProp name="Argument.value">${body}</stringProp>
                <stringProp name="Argument.metadata">=</stringProp>
              </elementProp>
            </collectionProp>
          </elementProp>
          <stringProp name="HTTPSampler.domain">hello-bi7moqtv2q-uc.a.run.app</stringProp>
          <stringProp name="HTTPSampler.port"></stringProp>
          <stringProp name="HTTPSampler.protocol">http</stringProp>
          <stringProp name="HTTPSampler.contentEncoding"></stringProp>
          <stringProp name="HTTPSampler.path">${url}</stringProp>
          <stringProp name="HTTPSampler.method">${method}</stringProp>
          <boolProp name="HTTPSampler.follow_redirects">true</boolProp>
          <boolProp name="HTTPSampler.auto_redirects">false</boolProp>
          <boolProp name="HTTPSampler.use_keepalive">true</boolProp>
          <boolProp name="HTTPSampler.DO_MULTIPART_POST">false</boolProp>
          <stringProp name="HTTPSampler.embedded_url_re"></stringProp>
          <stringProp name="HTTPSampler.connect_timeout"></stringProp>
          <stringProp name="HTTPSampler.response_timeout"></stringProp>
        </HTTPSamplerProxy>
        <hashTree>
          <HeaderManager guiclass="HeaderPanel" testclass="HeaderManager" testname="HTTP Header Manager" enabled="true">
            <collectionProp name="HeaderManager.headers">
              <elementProp name="" elementType="Header">
                <stringProp name="Header.name">User-Agent</stringProp>
                <stringProp name="Header.value">ApacheJMeter</stringProp>
              </elementProp>
              <elementProp name="" elementType="Header">
                <stringProp name="Header.name">Content-Type</stringProp>
                <stringProp name="Header.value">${contentType}</stringProp>
              </elementProp>
            </collectionProp>
          </HeaderManager>
          <hashTree/>
        </hashTree>
      </hashTree>
    </hashTree>
  </hashTree>
</jmeterTestPlan>
```

test-http.jmx

```
<stringProp name="HTTPSampler.domain">hello-bi7moqtv2q-uc.a.run.app</stringProp>
```

- 上述测试计划使用了 `ThreadGroup` 元素，其中设置了并发用户数为 `10`，循环执行 `5` 次。
- `HTTP Sampler` 配置了 GET 请求，目标 URL 为 `http://hello-bi7moqtv2q-uc.a.run.app`。
- 你可以通过修改 `<stringProp name="ThreadGroup.num_threads">` 来更改并发用户数，确保 `LoopController.loops` 设置为你想要的执行次数。

```
<?xml version="1.0" encoding="UTF-8"?>
<jmeterTestPlan version="1.2" properties="5.0" jmeter="5.4.1">
  <hashTree>
    <TestPlan guiclass="TestPlanGui" testclass="TestPlan" testname="Test Plan" enabled="true">
      <stringProp name="TestPlan.comments"></stringProp>
      <boolProp name="TestPlan.functional_mode">false</boolProp>
      <boolProp name="TestPlan.serialize_threadgroups">false</boolProp>
      <elementProp name="TestPlan.user_defined_variables" elementType="Arguments" guiclass="ArgumentsPanel" testclass="Arguments" testname="User Defined Variables" enabled="true">
        <collectionProp name="Arguments.arguments"/>
      </elementProp>
      <stringProp name="TestPlan.user_define_classpath"></stringProp>
    </TestPlan>
    <hashTree>
      <ThreadGroup guiclass="ThreadGroupGui" testclass="ThreadGroup" testname="Thread Group" enabled="true">
        <stringProp name="ThreadGroup.on_sample_error">continue</stringProp>
        <elementProp name="ThreadGroup.main_controller" elementType="LoopController" guiclass="LoopControllerGui" testclass="LoopController" testname="Loop Controller" enabled="true">
          <boolProp name="LoopController.continue_forever">false</boolProp>
          <stringProp name="LoopController.loops">5</stringProp>
        </elementProp>
        <stringProp name="ThreadGroup.num_threads">10</stringProp> <!-- 设置并发用户数 -->
        <stringProp name="ThreadGroup.ramp_time">1</stringProp>
        <boolProp name="ThreadGroup.scheduler">false</boolProp>
        <stringProp name="ThreadGroup.duration"></stringProp>
        <stringProp name="ThreadGroup.delay"></stringProp>
      </ThreadGroup>
      <hashTree>
        <HTTPSamplerProxy guiclass="HttpTestSampleGui" testclass="HTTPSamplerProxy" testname="HTTP Request" enabled="true">
          <boolProp name="HTTPSampler.postBodyRaw">true</boolProp>
          <elementProp name="HTTPsampler.Arguments" elementType="Arguments" guiclass="HTTPArgumentsPanel" testclass="Arguments" testname="User Defined Variables" enabled="true">
            <collectionProp name="Arguments.arguments"/>
          </elementProp>
          <stringProp name="HTTPSampler.domain">hello-bi7moqtv2q-uc.a.run.app</stringProp>
          <stringProp name="HTTPSampler.port"></stringProp>
          <stringProp name="HTTPSampler.protocol">http</stringProp>
          <stringProp name="HTTPSampler.contentEncoding"></stringProp>
          <stringProp name="HTTPSampler.path"></stringProp>
          <stringProp name="HTTPSampler.method">GET</stringProp>
          <boolProp name="HTTPSampler.follow_redirects">true</boolProp>
          <boolProp name="HTTPSampler.auto_redirects">false</boolProp>
          <boolProp name="HTTPSampler.use_keepalive">true</boolProp>
          <boolProp name="HTTPSampler.DO_MULTIPART_POST">false</boolProp>
          <boolProp name="HTTPSampler.monitor">false</boolProp>
          <stringProp name="HTTPSampler.embedded_url_re"></stringProp>
          <stringProp name="HTTPSampler.connect_timeout"></stringProp>
          <stringProp name="HTTPSampler.response_timeout"></stringProp>
        </HTTPSamplerProxy>
        <hashTree/>
      </hashTree>
    </hashTree>
  </hashTree>
</jmeterTestPlan>

```

test-https.jmx

```
<stringProp name="HTTPSampler.domain">hello-bi7moqtv2q-uc.a.run.app</stringProp>
```

`ThreadGroup.num_threads` 设置并发用户数，这里设置为 `10`。

`LoopController.loops` 设置循环次数，这里设置为 `5`。

`HTTPSampler.protocol` 设置为 `https`。

```
<?xml version="1.0" encoding="UTF-8"?>
<jmeterTestPlan version="1.2" properties="5.0" jmeter="5.4.1">
  <hashTree>
    <TestPlan guiclass="TestPlanGui" testclass="TestPlan" testname="Test Plan" enabled="true">
      <stringProp name="TestPlan.comments"></stringProp>
      <boolProp name="TestPlan.functional_mode">false</boolProp>
      <boolProp name="TestPlan.serialize_threadgroups">false</boolProp>
      <elementProp name="TestPlan.user_defined_variables" elementType="Arguments" guiclass="ArgumentsPanel" testclass="Arguments" testname="User Defined Variables" enabled="true">
        <collectionProp name="Arguments.arguments"/>
      </elementProp>
      <stringProp name="TestPlan.user_define_classpath"></stringProp>
    </TestPlan>
    <hashTree>
      <ThreadGroup guiclass="ThreadGroupGui" testclass="ThreadGroup" testname="Thread Group" enabled="true">
        <stringProp name="ThreadGroup.on_sample_error">continue</stringProp>
        <elementProp name="ThreadGroup.main_controller" elementType="LoopController" guiclass="LoopControllerGui" testclass="LoopController" testname="Loop Controller" enabled="true">
          <boolProp name="LoopController.continue_forever">false</boolProp>
          <stringProp name="LoopController.loops">50</stringProp>
        </elementProp>
        <stringProp name="ThreadGroup.num_threads">100</stringProp>
        <stringProp name="ThreadGroup.ramp_time">1</stringProp>
        <boolProp name="ThreadGroup.scheduler">false</boolProp>
        <stringProp name="ThreadGroup.duration"></stringProp>
        <stringProp name="ThreadGroup.delay"></stringProp>
      </ThreadGroup>
      <hashTree>
        <HTTPSamplerProxy guiclass="HttpTestSampleGui" testclass="HTTPSamplerProxy" testname="HTTP Request" enabled="true">
          <elementProp name="HTTPsampler.Arguments" elementType="Arguments" guiclass="HTTPArgumentsPanel" testclass="Arguments" testname="User Defined Variables" enabled="true">
            <collectionProp name="Arguments.arguments"/>
          </elementProp>
          <stringProp name="HTTPSampler.domain">hello-bi7moqtv2q-uc.a.run.app</stringProp>
          <stringProp name="HTTPSampler.port"></stringProp>
          <stringProp name="HTTPSampler.protocol">https</stringProp>
          <stringProp name="HTTPSampler.contentEncoding"></stringProp>
          <stringProp name="HTTPSampler.path"></stringProp>
          <stringProp name="HTTPSampler.method">GET</stringProp>
          <boolProp name="HTTPSampler.follow_redirects">true</boolProp>
          <boolProp name="HTTPSampler.auto_redirects">false</boolProp>
          <boolProp name="HTTPSampler.use_keepalive">true</boolProp>
          <boolProp name="HTTPSampler.DO_MULTIPART_POST">false</boolProp>
          <boolProp name="HTTPSampler.monitor">false</boolProp>
          <stringProp name="HTTPSampler.embedded_url_re"></stringProp>
          <stringProp name="HTTPSampler.connect_timeout"></stringProp>
          <stringProp name="HTTPSampler.response_timeout"></stringProp>
        </HTTPSamplerProxy>
        <hashTree/>
      </hashTree>
    </hashTree>
  </hashTree>
</jmeterTestPlan>

```

```
sudo  ./jmeter -n -t  templates/test-https.jmx -r
```





将 JMeter 测试计划设置为每秒钟启动 5 个用户，持续执行 5 分钟：

```
<?xml version="1.0" encoding="UTF-8"?>
<jmeterTestPlan version="1.2" properties="5.0" jmeter="5.4.1">
  <hashTree>
    <TestPlan guiclass="TestPlanGui" testclass="TestPlan" testname="Test Plan" enabled="true">
      <stringProp name="TestPlan.comments"></stringProp>
      <boolProp name="TestPlan.functional_mode">false</boolProp>
      <boolProp name="TestPlan.serialize_threadgroups">false</boolProp>
      <elementProp name="TestPlan.user_defined_variables" elementType="Arguments" guiclass="ArgumentsPanel" testclass="Arguments" testname="User Defined Variables" enabled="true">
        <collectionProp name="Arguments.arguments"/>
      </elementProp>
      <stringProp name="TestPlan.user_define_classpath"></stringProp>
    </TestPlan>
    <hashTree>
      <ThreadGroup guiclass="ThreadGroupGui" testclass="ThreadGroup" testname="Thread Group" enabled="true">
        <stringProp name="ThreadGroup.on_sample_error">continue</stringProp>
        <elementProp name="ThreadGroup.main_controller" elementType="LoopController" guiclass="LoopControllerGui" testclass="LoopController" testname="Loop Controller" enabled="true">
          <boolProp name="LoopController.continue_forever">false</boolProp>
          <stringProp name="LoopController.loops">-1</stringProp>
        </elementProp>
        <stringProp name="ThreadGroup.num_threads">5</stringProp>
        <stringProp name="ThreadGroup.ramp_time">1</stringProp>
        <boolProp name="ThreadGroup.scheduler">true</boolProp>
        <stringProp name="ThreadGroup.duration">300</stringProp>
        <stringProp name="ThreadGroup.delay"></stringProp>
      </ThreadGroup>
      <hashTree>
        <HTTPSamplerProxy guiclass="HttpTestSampleGui" testclass="HTTPSamplerProxy" testname="HTTP Request" enabled="true">
          <elementProp name="HTTPsampler.Arguments" elementType="Arguments" guiclass="HTTPArgumentsPanel" testclass="Arguments" testname="User Defined Variables" enabled="true">
            <collectionProp name="Arguments.arguments"/>
          </elementProp>
          <stringProp name="HTTPSampler.domain">hello-bi7moqtv2q-uc.a.run.app</stringProp>
          <stringProp name="HTTPSampler.port"></stringProp>
          <stringProp name="HTTPSampler.protocol">https</stringProp>
          <stringProp name="HTTPSampler.contentEncoding"></stringProp>
          <stringProp name="HTTPSampler.path"></stringProp>
          <stringProp name="HTTPSampler.method">GET</stringProp>
          <boolProp name="HTTPSampler.follow_redirects">true</boolProp>
          <boolProp name="HTTPSampler.auto_redirects">false</boolProp>
          <boolProp name="HTTPSampler.use_keepalive">true</boolProp>
          <boolProp name="HTTPSampler.DO_MULTIPART_POST">false</boolProp>
          <boolProp name="HTTPSampler.monitor">false</boolProp>
          <stringProp name="HTTPSampler.embedded_url_re"></stringProp>
          <stringProp name="HTTPSampler.connect_timeout"></stringProp>
          <stringProp name="HTTPSampler.response_timeout"></stringProp>
        </HTTPSamplerProxy>
        <hashTree/>
      </hashTree>
    </hashTree>
  </hashTree>
</jmeterTestPlan>

```









## 負荷テスト実施

- 負荷テストについて

負荷テストについては、全てmaster側で操作を行う
 今回は事前に負荷テスト用のjmxファイルを用意いたしました。
 `http://192.168.119.99/index.html` を5回叩く設定を`test.jmx`に書き込んでおります。
 jmxファイルはGUIで作成するのをおすすめします。

- テスト計画ファイルの作成の際に作ったテスト計画のファイルをjmeterサーバの**masterのみ**に設置

```
ll /usr/local/apache-jmeter-5.3/bin/templates/test.jmx 
```

- 実施

```
slaveのjmeterを使って、test.jmxファイルに記述したことをCUIモードで実行する

cd /usr/local/apache-jmeter-5.3/bin
./jmeter -n -t templates/test.jmx -r
----------
-n : CUIモードで実行
-t : シナリオファイル(ここで言うtest.jmx)を指定
-r : slaveのjmeterを利用する
---以下ここでは使っていないがよく使うオプション---
-l : 結果ファイルを指定
-e : レポートを作成
----------
cd /usr/local/apache-jmeter-5.3/bin/
./jmeter -n -t templates/test.jmx -r
==========
OpenJDK 64-Bit Server VM warning: If the number of processors is expected to increase from one, then you should configure the number of parallel GC threads appropriately using -XX:ParallelGCThreads=N
Creating summariser <summary>
Created the tree successfully using templates/test.jmx
Configuring remote engine: 192.168.119.101
Starting distributed test with remote engines: [192.168.119.101] @ Tue Sep 29 14:06:52 JST 2020 (1601356012593)
Remote engines have been started:[192.168.119.101]
Waiting for possible Shutdown/StopTestNow/HeapDump/ThreadDump message on port 4445
summary =      5 in 00:00:01 =    5.7/s Avg:     2 Min:     1 Max:     3 Err:     0 (0.00%)
Tidying up remote @ Tue Sep 29 14:06:54 JST 2020 (1601356014825)
... end of run
==========
```

- webサーバ側のlog

```
tail -f /var/log/httpd/access_log 
==========
192.168.119.101 - - [29/Sep/2020:14:06:34 +0900] "GET /index.html HTTP/1.1" 200 12 "-" "Apache-HttpClient/4.5.12 (Java/1.8.0_262)"
192.168.119.101 - - [29/Sep/2020:14:06:34 +0900] "GET /index.html HTTP/1.1" 200 12 "-" "Apache-HttpClient/4.5.12 (Java/1.8.0_262)"
192.168.119.101 - - [29/Sep/2020:14:06:34 +0900] "GET /index.html HTTP/1.1" 200 12 "-" "Apache-HttpClient/4.5.12 (Java/1.8.0_262)"
192.168.119.101 - - [29/Sep/2020:14:06:54 +0900] "GET /index.html HTTP/1.1" 200 12 "-" "Apache-HttpClient/4.5.12 (Java/1.8.0_262)"
192.168.119.101 - - [29/Sep/2020:14:06:54 +0900] "GET /index.html HTTP/1.1" 200 12 "-" "Apache-HttpClient/4.5.12 (Java/1.8.0_262)"
192.168.119.101 - - [29/Sep/2020:14:06:54 +0900] "GET /index.html HTTP/1.1" 200 12 "-" "Apache-HttpClient/4.5.12 (Java/1.8.0_262)"
192.168.119.101 - - [29/Sep/2020:14:06:54 +0900] "GET /index.html HTTP/1.1" 200 12 "-" "Apache-HttpClient/4.5.12 (Java/1.8.0_262)"
192.168.119.101 - - [29/Sep/2020:14:06:54 +0900] "GET /index.html HTTP/1.1" 200 12 "-" "Apache-HttpClient/4.5.12 (Java/1.8.0_262)"
==========
slaveのjmeterから5回、index.htmlにアクセスがあった
```

## 